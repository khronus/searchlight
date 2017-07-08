package fake

import (
	aci "github.com/appscode/searchlight/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"
)

type FakeAlert struct {
	Fake *testing.Fake
	ns   string
}

var resourceAlert = schema.GroupVersionResource{Group: "monitoring.appscode.com", Version: "v1alpha1", Resource: "alerts"}

// Get returns the Alert by name.
func (mock *FakeAlert) Get(name string) (*aci.Alert, error) {
	obj, err := mock.Fake.
		Invokes(testing.NewGetAction(resourceAlert, mock.ns, name), &aci.Alert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aci.Alert), err
}

// List returns the a of Alerts.
func (mock *FakeAlert) List(opts metav1.ListOptions) (*aci.AlertList, error) {
	obj, err := mock.Fake.
		Invokes(testing.NewListAction(resourceAlert, mock.ns, opts), &aci.Alert{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &aci.AlertList{}
	for _, item := range obj.(*aci.AlertList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Create creates a new Alert.
func (mock *FakeAlert) Create(svc *aci.Alert) (*aci.Alert, error) {
	obj, err := mock.Fake.
		Invokes(testing.NewCreateAction(resourceAlert, mock.ns, svc), &aci.Alert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aci.Alert), err
}

// Update updates a Alert.
func (mock *FakeAlert) Update(svc *aci.Alert) (*aci.Alert, error) {
	obj, err := mock.Fake.
		Invokes(testing.NewUpdateAction(resourceAlert, mock.ns, svc), &aci.Alert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aci.Alert), err
}

// Delete deletes a Alert by name.
func (mock *FakeAlert) Delete(name string) error {
	_, err := mock.Fake.
		Invokes(testing.NewDeleteAction(resourceAlert, mock.ns, name), &aci.Alert{})

	return err
}

func (mock *FakeAlert) UpdateStatus(srv *aci.Alert) (*aci.Alert, error) {
	obj, err := mock.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resourceAlert, "status", mock.ns, srv), &aci.Alert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*aci.Alert), err
}

func (mock *FakeAlert) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return mock.Fake.
		InvokesWatch(testing.NewWatchAction(resourceAlert, mock.ns, opts))
}
