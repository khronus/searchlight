package plugin

import (
	"encoding/json"
	"sync"

	api "github.com/appscode/searchlight/apis/monitoring/v1alpha1"
	admission "k8s.io/api/admission/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	hooks "kmodules.xyz/webhook-runtime/admission/v1beta1"
)

type CRDValidator struct {
	client      kubernetes.Interface
	lock        sync.RWMutex
	initialized bool
}

func (a *CRDValidator) Resource() (plural schema.GroupVersionResource, singular string) {
	return schema.GroupVersionResource{
			Group:    "admission.monitoring.appscode.com",
			Version:  "v1alpha1",
			Resource: "admissionreviews",
		},
		"admissionreview"
}

func (a *CRDValidator) Initialize(config *rest.Config, stopCh <-chan struct{}) error {
	a.lock.Lock()
	defer a.lock.Unlock()

	a.initialized = true

	var err error
	if a.client, err = kubernetes.NewForConfig(config); err != nil {
		return err
	}
	return err
}

func (a *CRDValidator) Admit(req *admission.AdmissionRequest) *admission.AdmissionResponse {
	status := &admission.AdmissionResponse{}
	supportedKinds := sets.NewString(api.ResourceKindClusterAlert, api.ResourceKindNodeAlert, api.ResourceKindPodAlert)

	if (req.Operation != admission.Create && req.Operation != admission.Update) ||
		len(req.SubResource) != 0 ||
		req.Kind.Group != api.SchemeGroupVersion.Group ||
		!supportedKinds.Has(req.Kind.Kind) {
		status.Allowed = true
		return status
	}

	a.lock.RLock()
	defer a.lock.RUnlock()
	if !a.initialized {
		return hooks.StatusUninitialized()
	}

	var alert api.Alert
	switch req.Kind.Kind {
	case api.ResourceKindClusterAlert:
		alert = &api.ClusterAlert{}
	case api.ResourceKindNodeAlert:
		alert = &api.NodeAlert{}
	case api.ResourceKindPodAlert:
		alert = &api.PodAlert{}
	}

	err := json.Unmarshal(req.Object.Raw, alert)
	if err != nil {
		return hooks.StatusBadRequest(err)
	}
	err = alert.IsValid(a.client)
	if err != nil {
		return hooks.StatusForbidden(err)
	}

	status.Allowed = true
	return status
}
