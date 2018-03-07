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
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// nodeAlertMapperConfiguration
type naMapperConf struct {
	Selector map[string]string `json:"selector,omitempty"`
	NodeName string            `json:"nodeName,omitempty"`
}

func (op *Operator) initNodeAlertWatcher() {
	op.naInformer = op.monInformerFactory.Monitoring().V1alpha1().NodeAlerts().Informer()
	op.naQueue = queue.New("NodeAlert", op.options.MaxNumRequeues, op.options.NumThreads, op.reconcileNodeAlert)
	op.naInformer.AddEventHandler(&cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			alert := obj.(*api.NodeAlert)
			if op.validateAlert(alert) {
				queue.Enqueue(op.caQueue.GetQueue(), obj)
			}
		},
		UpdateFunc: func(oldObj interface{}, newObj interface{}) {
			old := oldObj.(*api.NodeAlert)
			nu := newObj.(*api.NodeAlert)

			// DeepEqual old & new
			// DeepEqual MapperConfiguration of old & new
			// Patch PodAlert with necessary annotation
			nu, proceed, err := op.processNodeAlertUpdate(old, nu)
			if err != nil {
				log.Error(err)
			} else if proceed {
				queue.Enqueue(op.naQueue.GetQueue(), nu)
			}
		},
		DeleteFunc: func(obj interface{}) {
			queue.Enqueue(op.naQueue.GetQueue(), obj)
		},
	})
	op.naLister = op.monInformerFactory.Monitoring().V1alpha1().NodeAlerts().Lister()
}

func (op *Operator) reconcileNodeAlert(key string) error {
	obj, exists, err := op.naInformer.GetIndexer().GetByKey(key)
	if err != nil {
		glog.Errorf("Fetching object with key %s from store failed with %v", key, err)
		return err
	}
	if !exists {
		log.Debugf("NodeAlert %s does not exist anymore\n", key)
		return nil
	}

	alert := obj.(*api.NodeAlert)
	if alert.DeletionTimestamp != nil {
		if core_util.HasFinalizer(alert.ObjectMeta, mon_api.GroupName) {
			// Delete all Icinga objects created for this NodeAlert
			if err := op.EnsureNodeAlertDeleted(alert); err != nil {
				log.Errorf("Failed to delete NodeAlert %s@%s. Reason: %v", alert.Name, alert.Namespace, err)
				return err
			}
			// Remove Finalizer
			_, _, err = mon_util.PatchNodeAlert(op.ExtClient.MonitoringV1alpha1(), alert, func(in *api.NodeAlert) *api.NodeAlert {
				in.ObjectMeta = core_util.RemoveFinalizer(in.ObjectMeta, mon_api.GroupName)
				return in
			})
			return err
		}
	} else {
		log.Infof("Sync/Add/Update for NodeAlert %s\n", alert.GetName())

		alert, _, err = mon_util.PatchNodeAlert(op.ExtClient.MonitoringV1alpha1(), alert, func(in *api.NodeAlert) *api.NodeAlert {
			in.ObjectMeta = core_util.AddFinalizer(in.ObjectMeta, mon_api.GroupName)
			return in
		})

		if err := op.EnsureNodeAlert(alert); err != nil {
			log.Errorf("Failed to sync NodeAlert %s@%s. Reason: %v", alert.Name, alert.Namespace, err)
		}
	}
	return nil
}

func (op *Operator) EnsureNodeAlert(alert *api.NodeAlert) error {

	var oldMc *naMapperConf
	if val, ok := alert.Annotations[annotationLastConfiguration]; ok {
		if err := json.Unmarshal([]byte(val), &oldMc); err != nil {
			return err
		}
	}

	oldMappedNode, err := op.getMappedNodeList(alert.Namespace, oldMc)
	if err != nil {
		return err
	}

	newMC := &naMapperConf{
		Selector: alert.Spec.Selector,
		NodeName: alert.Spec.NodeName,
	}
	newMappedNode, err := op.getMappedNodeList(alert.Namespace, newMC)
	if err != nil {
		return err
	}

	for key, node := range newMappedNode {
		delete(oldMappedNode, node.Name)

		op.setNodeAlertNamesInAnnotation(node, alert)

		go op.EnsureIcingaNodeAlert(alert, newMappedNode[key])
	}

	for _, node := range oldMappedNode {
		op.EnsureIcingaNodeAlertDeleted(alert, node)
	}

	return nil
}

func (op *Operator) EnsureNodeAlertDeleted(alert *api.NodeAlert) error {
	mc := &naMapperConf{
		Selector: alert.Spec.Selector,
		NodeName: alert.Spec.NodeName,
	}
	mappedPod, err := op.getMappedNodeList(alert.Namespace, mc)
	if err != nil {
		return err
	}

	for _, node := range mappedPod {
		op.EnsureIcingaNodeAlertDeleted(alert, node)
	}

	return nil
}

func (op *Operator) EnsureIcingaNodeAlert(alert *api.NodeAlert, node *core.Node) (err error) {
	err = op.nodeHost.Reconcile(alert.DeepCopy(), node.DeepCopy())
	if err != nil {
		op.recorder.Eventf(
			alert.ObjectReference(),
			core.EventTypeWarning,
			eventer.EventReasonFailedToSync,
			`Reason: %v`,
			err,
		)
	}
	return
}

func (op *Operator) EnsureIcingaNodeAlertDeleted(alert *api.NodeAlert, node *core.Node) (err error) {
	err = op.nodeHost.Delete(alert, node)
	if err != nil && alert != nil {
		op.recorder.Eventf(
			alert.ObjectReference(),
			core.EventTypeWarning,
			eventer.EventReasonFailedToDelete,
			`Reason: %v`,
			err,
		)
	}
	return
}

func (op *Operator) processNodeAlertUpdate(old, nu *api.NodeAlert) (*api.NodeAlert, bool, error) {
	// Check for changes in Spec
	if !reflect.DeepEqual(old.Spec, nu.Spec) {
		if !op.validateAlert(nu) {
			return nil, false, errors.Errorf(`Invalid NodeAlert "%s@%s"`, nu.Name, nu.Namespace)
		}

		// We need Selector/NodeName from oldAlert while processing this update operation.
		// Because we need to remove Icinga objects for oldAlert.
		oldMC := &naMapperConf{
			Selector: old.Spec.Selector,
			NodeName: old.Spec.NodeName,
		}
		newMC := &naMapperConf{
			Selector: nu.Spec.Selector,
			NodeName: nu.Spec.NodeName,
		}

		// We will store Selector/PodName from oldAlert in annotation
		if !reflect.DeepEqual(oldMC, newMC) {
			var err error
			// Patch NodeAlert with Selector/PodName from oldAlert (oldMC)
			nu, _, err = mon_util.PatchNodeAlert(op.ExtClient.MonitoringV1alpha1(), nu, func(in *api.NodeAlert) *api.NodeAlert {
				if in.Annotations == nil {
					in.Annotations = make(map[string]string, 0)
				}
				data, _ := json.Marshal(oldMC)
				in.Annotations[annotationLastConfiguration] = string(data)
				return in
			})
			if err != nil {
				op.recorder.Eventf(
					nu.ObjectReference(),
					core.EventTypeWarning,
					eventer.EventReasonFailedToUpdate,
					`Reason: %v`,
					err,
				)
				return nil, false, errors.Wrap(err,
					fmt.Sprintf(`Failed to patch PodAlert "%s@%s"`, nu.Name, nu.Namespace),
				)
			}
		}
		return nu, true, nil
	} else if nu.DeletionTimestamp != nil {
		return nu, true, nil
	}

	return nu, false, nil
}

func (op *Operator) getMappedNodeList(namespace string, mc *naMapperConf) (map[string]*core.Node, error) {
	mappedPodList := make(map[string]*core.Node)

	if mc == nil {
		return mappedPodList, nil
	}

	sel := labels.SelectorFromSet(mc.Selector)

	if mc.NodeName != "" {
		if node, err := op.nLister.Get(mc.NodeName); err == nil {
			if sel.Matches(labels.Set(node.Labels)) {
				mappedPodList[node.Name] = node
			}
		}
	} else {
		if nodeList, err := op.nLister.List(sel); err != nil {
			return nil, err
		} else {
			for i := range nodeList {
				node := nodeList[i]
				mappedPodList[node.Name] = node
			}
		}
	}

	return mappedPodList, nil
}

func (op *Operator) setNodeAlertNamesInAnnotation(node *core.Node, alert *api.NodeAlert) {
	_, _, err := core_util.PatchNode(op.KubeClient, node, func(in *core.Node) *core.Node {
		if in.Annotations == nil {
			in.Annotations = make(map[string]string, 0)
		}

		alertNames := make([]string, 0)
		if val, ok := alert.Annotations[annotationAlertsName]; ok {
			if err := json.Unmarshal([]byte(val), &alertNames); err != nil {
				log.Errorf("Failed to patch Node %s.", node.Name)
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
		log.Errorf("Failed to patch Node %s.", node.Name)
	}
}
