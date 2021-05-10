/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	monitoringv1alpha1 "go.searchlight.dev/icinga-operator/apis/monitoring/v1alpha1"
	versioned "go.searchlight.dev/icinga-operator/client/clientset/versioned"
	internalinterfaces "go.searchlight.dev/icinga-operator/client/informers/externalversions/internalinterfaces"
	v1alpha1 "go.searchlight.dev/icinga-operator/client/listers/monitoring/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PodAlertInformer provides access to a shared informer and lister for
// PodAlerts.
type PodAlertInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PodAlertLister
}

type podAlertInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPodAlertInformer constructs a new informer for PodAlert type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodAlertInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodAlertInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPodAlertInformer constructs a new informer for PodAlert type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodAlertInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1alpha1().PodAlerts(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MonitoringV1alpha1().PodAlerts(namespace).Watch(context.TODO(), options)
			},
		},
		&monitoringv1alpha1.PodAlert{},
		resyncPeriod,
		indexers,
	)
}

func (f *podAlertInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodAlertInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podAlertInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&monitoringv1alpha1.PodAlert{}, f.defaultInformer)
}

func (f *podAlertInformer) Lister() v1alpha1.PodAlertLister {
	return v1alpha1.NewPodAlertLister(f.Informer().GetIndexer())
}
