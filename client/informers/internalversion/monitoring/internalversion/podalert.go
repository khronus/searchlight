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

package internalversion

import (
	time "time"

	monitoring "github.com/appscode/searchlight/apis/monitoring"
	clientset_internalversion "github.com/appscode/searchlight/client/clientset/internalversion"
	internalinterfaces "github.com/appscode/searchlight/client/informers/internalversion/internalinterfaces"
	internalversion "github.com/appscode/searchlight/client/listers/monitoring/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PodAlertInformer provides access to a shared informer and lister for
// PodAlerts.
type PodAlertInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.PodAlertLister
}

type podAlertInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPodAlertInformer constructs a new informer for PodAlert type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodAlertInformer(client clientset_internalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodAlertInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPodAlertInformer constructs a new informer for PodAlert type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodAlertInformer(client clientset_internalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Monitoring().PodAlerts(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Monitoring().PodAlerts(namespace).Watch(options)
			},
		},
		&monitoring.PodAlert{},
		resyncPeriod,
		indexers,
	)
}

func (f *podAlertInformer) defaultInformer(client clientset_internalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodAlertInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podAlertInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&monitoring.PodAlert{}, f.defaultInformer)
}

func (f *podAlertInformer) Lister() internalversion.PodAlertLister {
	return internalversion.NewPodAlertLister(f.Informer().GetIndexer())
}
