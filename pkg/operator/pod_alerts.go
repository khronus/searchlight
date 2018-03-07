package operator

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/appscode/go/log"
	"github.com/appscode/go/sets"
	core_util "github.com/appscode/kutil/core/v1"
	"github.com/appscode/kutil/tools/queue"
	mon_api "github.com/appscode/searchlight/apis/monitoring"
	api "github.com/appscode/searchlight/apis/monitoring/v1alpha1"
	mon_util "github.com/appscode/searchlight/client/clientset/versioned/typed/monitoring/v1alpha1/util"
	"github.com/appscode/searchlight/pkg/eventer"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PodAlertMapperConfiguration
type paMapperConf struct {
	Selector metav1.LabelSelector `json:"selector,omitempty"`
	PodName  string               `json:"podName,omitempty"`
}

func (op *Operator) initPodAlertWatcher() {
	op.paInformer = op.monInformerFactory.Monitoring().V1alpha1().PodAlerts().Informer()
	op.paQueue = queue.New("PodAlert", op.options.MaxNumRequeues, op.options.NumThreads, op.reconcilePodAlert)
	op.paInformer.AddEventHandler(&cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			alert := obj.(*api.PodAlert)
			if op.validateAlert(alert) {
				queue.Enqueue(op.paQueue.GetQueue(), obj)
			}
		},
		UpdateFunc: func(oldObj interface{}, newObj interface{}) {
			old := oldObj.(*api.PodAlert)
			nu := newObj.(*api.PodAlert)

			// DeepEqual old & new
			// DeepEqual MapperConfiguration of old & new
			// Patch PodAlert with necessary annotation
			nu, proceed, err := op.processPodAlertUpdate(old, nu)
			if err != nil {
				log.Error(err)
			} else if proceed {
				queue.Enqueue(op.paQueue.GetQueue(), nu)
			}
		},
		DeleteFunc: func(obj interface{}) {
			queue.Enqueue(op.paQueue.GetQueue(), obj)
		},
	})
	op.paLister = op.monInformerFactory.Monitoring().V1alpha1().PodAlerts().Lister()
}

func (op *Operator) reconcilePodAlert(key string) error {
	obj, exists, err := op.paInformer.GetIndexer().GetByKey(key)
	if err != nil {
		glog.Errorf("Fetching object with key %s from store failed with %v", key, err)
		return err
	}

	if !exists {
		log.Debugf("PodAlert %s does not exist anymore\n", key)
	} else {
		alert := obj.(*api.PodAlert)
		if alert.DeletionTimestamp != nil {
			if core_util.HasFinalizer(alert.ObjectMeta, mon_api.GroupName) {
				// Delete all Icinga objects created for this PodAlert
				if err := op.EnsurePodAlertDeleted(alert); err != nil {
					log.Errorf("Failed to delete PodAlert %s@%s. Reason: %v", alert.Name, alert.Namespace, err)
					return err
				}
				// Remove Finalizer
				_, _, err = mon_util.PatchPodAlert(op.ExtClient.MonitoringV1alpha1(), alert, func(in *api.PodAlert) *api.PodAlert {
					in.ObjectMeta = core_util.RemoveFinalizer(in.ObjectMeta, mon_api.GroupName)
					return in
				})
				return err
			}
		} else {
			log.Infof("Sync/Add/Update for PodAlert %s\n", alert.GetName())

			// Patch PodAlert to add Finalizer
			alert, _, _ = mon_util.PatchPodAlert(op.ExtClient.MonitoringV1alpha1(), alert, func(in *api.PodAlert) *api.PodAlert {
				in.ObjectMeta = core_util.AddFinalizer(in.ObjectMeta, mon_api.GroupName)
				return in
			})

			if err := op.EnsurePodAlert(alert); err != nil {
				log.Errorf("Failed to sync PodAlert %s@%s. Reason: %v", alert.Name, alert.Namespace, err)
				return err
			}
		}
	}
	return nil
}

func (op *Operator) EnsurePodAlert(alert *api.PodAlert) error {

	var oldMc *paMapperConf
	if val, ok := alert.Annotations[annotationLastConfiguration]; ok {
		if err := json.Unmarshal([]byte(val), &oldMc); err != nil {
			return err
		}
	}

	oldMappedPod, err := op.getMappedPodList(alert.Namespace, oldMc)
	if err != nil {
		return err
	}

	newMC := &paMapperConf{
		Selector: alert.Spec.Selector,
		PodName:  alert.Spec.PodName,
	}
	newMappedPod, err := op.getMappedPodList(alert.Namespace, newMC)
	if err != nil {
		return err
	}

	for key, pod := range newMappedPod {
		delete(oldMappedPod, pod.Name)
		if pod.Status.PodIP == "" {
			log.Warningf("Skipping pod %s@%s, since it has no IP", pod.Name, pod.Namespace)
		}

		op.setPodAlertNamesInAnnotation(pod, alert)

		go op.EnsureIcingaPodAlert(alert, newMappedPod[key])
	}

	for _, pod := range oldMappedPod {
		op.EnsureIcingaPodAlertDeleted(alert, pod)
	}

	return nil
}

func (op *Operator) EnsurePodAlertDeleted(alert *api.PodAlert) error {
	mc := &paMapperConf{
		Selector: alert.Spec.Selector,
		PodName:  alert.Spec.PodName,
	}
	mappedPod, err := op.getMappedPodList(alert.Namespace, mc)
	if err != nil {
		return err
	}

	for _, pod := range mappedPod {
		op.EnsureIcingaPodAlertDeleted(alert, pod)
	}

	return nil
}

func (op *Operator) EnsureIcingaPodAlert(alert *api.PodAlert, pod *core.Pod) (err error) {
	err = op.podHost.Reconcile(alert.DeepCopy(), pod.DeepCopy())
	if err != nil {
		op.recorder.Eventf(
			alert.ObjectReference(),
			core.EventTypeWarning,
			eventer.EventReasonFailedToSync,
			`Reason: %v`,
			alert.Name,
			err,
		)
	}
	return err
}

func (op *Operator) EnsureIcingaPodAlertDeleted(alert *api.PodAlert, pod *core.Pod) (err error) {
	err = op.podHost.Delete(alert, pod)
	if err != nil && alert != nil {
		op.recorder.Eventf(
			alert.ObjectReference(),
			core.EventTypeWarning,
			eventer.EventReasonFailedToDelete,
			`Fail to delete Icinga objects of PodAlert "%s@%s" for Pod "%s@%s". Reason: %v`,
			alert.Name, alert.Namespace, pod.Name, pod.Namespace,
			err,
		)
	}
	return err
}

func (op *Operator) processPodAlertUpdate(oldAlert, newAlert *api.PodAlert) (*api.PodAlert, bool, error) {
	// Check for changes in Spec
	if !reflect.DeepEqual(oldAlert.Spec, newAlert.Spec) {
		if !op.validateAlert(newAlert) {
			return nil, false, errors.Errorf(`Invalid PodAlert "%s@%s"`, newAlert.Name, newAlert.Namespace)
		}

		// We need Selector/PodName from oldAlert while processing this update operation.
		// Because we need to remove Icinga objects for oldAlert.
		oldMC := &paMapperConf{
			Selector: oldAlert.Spec.Selector,
			PodName:  oldAlert.Spec.PodName,
		}
		newMC := &paMapperConf{
			Selector: newAlert.Spec.Selector,
			PodName:  newAlert.Spec.PodName,
		}

		// We will store Selector/PodName from oldAlert in annotation
		if !reflect.DeepEqual(oldMC, newMC) {
			var err error
			// Patch PodAlert with Selector/PodName from oldAlert (oldMC)
			newAlert, _, err = mon_util.PatchPodAlert(
				op.ExtClient.MonitoringV1alpha1(),
				newAlert,
				func(in *api.PodAlert) *api.PodAlert {
					if in.Annotations == nil {
						in.Annotations = make(map[string]string, 0)
					}
					data, _ := json.Marshal(oldMC)
					in.Annotations[annotationLastConfiguration] = string(data)
					return in
				})
			if err != nil {
				op.recorder.Eventf(
					newAlert.ObjectReference(),
					core.EventTypeWarning,
					eventer.EventReasonFailedToUpdate,
					`Reason: %v`,
					err,
				)
				return nil, false, errors.Wrap(err,
					fmt.Sprintf(`Failed to patch PodAlert "%s@%s"`, newAlert.Name, newAlert.Namespace),
				)
			}
		}
		return newAlert, true, nil
	} else if newAlert.DeletionTimestamp != nil {
		return newAlert, true, nil
	}

	return newAlert, false, nil
}

func (op *Operator) getMappedPodList(namespace string, mc *paMapperConf) (map[string]*core.Pod, error) {
	mappedPodList := make(map[string]*core.Pod)

	if mc == nil {
		return mappedPodList, nil
	}

	sel, err := metav1.LabelSelectorAsSelector(&mc.Selector)
	if err != nil {
		return nil, err
	}
	if mc.PodName != "" {
		if pod, err := op.pLister.Pods(namespace).Get(mc.PodName); err == nil {
			if sel.Matches(labels.Set(pod.Labels)) {
				mappedPodList[pod.Name] = pod
			}
		}
	} else {
		if podList, err := op.pLister.Pods(namespace).List(sel); err != nil {
			return nil, err
		} else {
			for i := range podList {
				pod := podList[i]
				mappedPodList[pod.Name] = pod
			}
		}
	}

	return mappedPodList, nil
}

func (op *Operator) setPodAlertNamesInAnnotation(pod *core.Pod, alert *api.PodAlert) {
	_, _, err := core_util.PatchPod(op.KubeClient, pod, func(in *core.Pod) *core.Pod {
		if in.Annotations == nil {
			in.Annotations = make(map[string]string, 0)
		}

		alertNames := make([]string, 0)
		if val, ok := pod.Annotations[annotationAlertsName]; ok {
			if err := json.Unmarshal([]byte(val), &alertNames); err != nil {
				log.Errorf("Failed to patch Pod %s@%s.", pod.Name, pod.Namespace)
			}
		}
		ss := sets.NewString(alertNames...)
		ss.Insert(alert.Name)
		alertNames = ss.List()
		data, _ := json.Marshal(alertNames)
		in.Annotations[annotationAlertsName] = string(data)
		return in
	})

	if err != nil {
		log.Errorf("Failed to patch Pod %s@%s.", pod.Name, pod.Namespace)
	}
}
