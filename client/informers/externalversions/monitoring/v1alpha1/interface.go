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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	internalinterfaces "github.com/appscode/searchlight/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterAlerts returns a ClusterAlertInformer.
	ClusterAlerts() ClusterAlertInformer
	// Incidents returns a IncidentInformer.
	Incidents() IncidentInformer
	// NodeAlerts returns a NodeAlertInformer.
	NodeAlerts() NodeAlertInformer
	// PodAlerts returns a PodAlertInformer.
	PodAlerts() PodAlertInformer
	// SearchlightPlugins returns a SearchlightPluginInformer.
	SearchlightPlugins() SearchlightPluginInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ClusterAlerts returns a ClusterAlertInformer.
func (v *version) ClusterAlerts() ClusterAlertInformer {
	return &clusterAlertInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Incidents returns a IncidentInformer.
func (v *version) Incidents() IncidentInformer {
	return &incidentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NodeAlerts returns a NodeAlertInformer.
func (v *version) NodeAlerts() NodeAlertInformer {
	return &nodeAlertInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PodAlerts returns a PodAlertInformer.
func (v *version) PodAlerts() PodAlertInformer {
	return &podAlertInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SearchlightPlugins returns a SearchlightPluginInformer.
func (v *version) SearchlightPlugins() SearchlightPluginInformer {
	return &searchlightPluginInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
