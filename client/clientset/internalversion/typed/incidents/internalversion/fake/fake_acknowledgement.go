/*
Copyright 2018 The Searchlight Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package fake

import (
	incidents "github.com/appscode/searchlight/apis/incidents"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// FakeAcknowledgements implements AcknowledgementInterface
type FakeAcknowledgements struct {
	Fake *FakeIncidents
	ns   string
}

var acknowledgementsResource = schema.GroupVersionResource{Group: "incidents.monitoring.appscode.com", Version: "", Resource: "acknowledgements"}

var acknowledgementsKind = schema.GroupVersionKind{Group: "incidents.monitoring.appscode.com", Version: "", Kind: "Acknowledgement"}

// Create takes the representation of a acknowledgement and creates it.  Returns the server's representation of the acknowledgement, and an error, if there is any.
func (c *FakeAcknowledgements) Create(acknowledgement *incidents.Acknowledgement) (result *incidents.Acknowledgement, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(acknowledgementsResource, c.ns, acknowledgement), &incidents.Acknowledgement{})

	if obj == nil {
		return nil, err
	}
	return obj.(*incidents.Acknowledgement), err
}

// Delete takes name of the acknowledgement and deletes it. Returns an error if one occurs.
func (c *FakeAcknowledgements) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(acknowledgementsResource, c.ns, name), &incidents.Acknowledgement{})

	return err
}