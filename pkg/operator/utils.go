package operator

import (
	api "github.com/appscode/searchlight/apis/monitoring/v1alpha1"
	mon_listers "github.com/appscode/searchlight/client/listers/monitoring/v1alpha1"
	"github.com/appscode/searchlight/pkg/eventer"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

func (op *Operator) isValid(alert api.Alert) bool {
	// Validate IcingaCommand & it's variables.
	// And also check supported IcingaState
	err := alert.IsValid(op.kubeClient)
	if err != nil {
		op.recorder.Eventf(
			alert.ObjectReference(),
			core.EventTypeWarning,
			eventer.EventReasonAlertInvalid,
			`Reason: %v`,
			err,
		)
	}
	return err == nil
}

func findPodAlert(kc kubernetes.Interface, lister mon_listers.PodAlertLister, obj metav1.ObjectMeta) ([]*api.PodAlert, error) {
	alerts, err := lister.PodAlerts(obj.Namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	result := make([]*api.PodAlert, 0)
	for i := range alerts {
		alert := alerts[i]
		if err := alert.IsValid(kc); err != nil {
			continue
		}

		if alert.Spec.PodName != nil {
			if *alert.Spec.PodName == obj.Name {
				result = append(result, alert)
			}
		} else if alert.Spec.Selector != nil {
			if selector, err := metav1.LabelSelectorAsSelector(alert.Spec.Selector); err == nil {
				if selector.Matches(labels.Set(obj.Labels)) {
					result = append(result, alert)
				}
			}
		}
	}
	return result, nil
}

func findNodeAlert(kc kubernetes.Interface, lister mon_listers.NodeAlertLister, obj metav1.ObjectMeta) ([]*api.NodeAlert, error) {
	alerts, err := lister.NodeAlerts(obj.Namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	result := make([]*api.NodeAlert, 0)
	for i := range alerts {
		alert := alerts[i]
		if err := alert.IsValid(kc); err != nil {
			continue
		}

		if alert.Spec.NodeName != nil {
			if *alert.Spec.NodeName == obj.Name {
				result = append(result, alert)
			}
		} else {
			selector := labels.SelectorFromSet(alert.Spec.Selector)
			if selector.Matches(labels.Set(obj.Labels)) {
				result = append(result, alert)
			}
		}
	}
	return result, nil
}
