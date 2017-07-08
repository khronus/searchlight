package watcher

import (
	"reflect"
	"sync"
	"time"

	"github.com/appscode/log"
	aci "github.com/appscode/searchlight/api"
	acs "github.com/appscode/searchlight/client/clientset"
	"github.com/appscode/searchlight/pkg/analytics"
	"github.com/appscode/searchlight/pkg/client/icinga"
	"github.com/appscode/searchlight/pkg/controller"
	"github.com/appscode/searchlight/pkg/events"
	"github.com/appscode/searchlight/pkg/stash"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	clientset "k8s.io/client-go/kubernetes"
	apiv1 "k8s.io/client-go/pkg/api/v1"
	extensions "k8s.io/client-go/pkg/apis/extensions/v1beta1"
	"k8s.io/client-go/tools/cache"
)

type Watcher struct {
	KubeClient   clientset.Interface
	ExtClient    acs.ExtensionInterface
	IcingaClient *icinga.IcingaClient

	// sync time to sync the list.
	SyncPeriod time.Duration

	// lister store
	Storage stash.Storage
	// Enable analytics
	EnableAnalytics bool
	sync.Mutex
}

func (w *Watcher) Run() {
	w.setup()
	w.Service()
	w.StatefulSet()
	w.DaemonSet()
	w.ReplicaSet()
	w.RC()
	w.Pod()
	w.Alert()
	w.AlertEvent()
	w.Node()
	w.Deployment()
}

func (w *Watcher) setup() {
	log.Infoln("Ensuring ThirdPartyResource")

	if err := w.ensureThirdPartyResource(aci.ResourceNameAlert + "." + aci.V1alpha1SchemeGroupVersion.Group); err != nil {
		log.Fatalln(err)
	}
	if err := w.ensureThirdPartyResource(aci.ResourceNameNodeAlert + "." + aci.V1alpha1SchemeGroupVersion.Group); err != nil {
		log.Fatalln(err)
	}
	if err := w.ensureThirdPartyResource(aci.ResourceNameClusterAlert + "." + aci.V1alpha1SchemeGroupVersion.Group); err != nil {
		log.Fatalln(err)
	}

	// Enable analytics
	if w.EnableAnalytics {
		analytics.Enable()
	}
}

func (w *Watcher) ensureThirdPartyResource(resourceName string) error {
	_, err := w.KubeClient.ExtensionsV1beta1().ThirdPartyResources().Get(resourceName, metav1.GetOptions{})
	if !kerr.IsNotFound(err) {
		return err
	}

	thirdPartyResource := &extensions.ThirdPartyResource{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "extensions/v1beta1",
			Kind:       "ThirdPartyResource",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: resourceName,
			Labels: map[string]string{
				"app": "searchlight",
			},
		},
		Description: "Searchlight by AppsCode - Alerts for Kubernetes",
		Versions: []extensions.APIVersion{
			{
				Name: aci.V1alpha1SchemeGroupVersion.Version,
			},
		},
	}

	_, err = w.KubeClient.ExtensionsV1beta1().ThirdPartyResources().Create(thirdPartyResource)
	return err
}

func (w *Watcher) Dispatch(e *events.Event) error {
	if e.Ignorable() {
		return nil
	}
	log.Debugln("Dispatching event with resource", e.ResourceType, "event", e.EventType)
	if e.ResourceType == events.Alert || e.ResourceType == events.Node ||
		e.ResourceType == events.Pod || e.ResourceType == events.Service || e.ResourceType == events.AlertEvent {
		return controller.New(w.KubeClient, w.IcingaClient, w.ExtClient, w.Storage).Handle(e)
	}
	return nil
}

func (w *Watcher) Cache(resource events.ObjectType, object runtime.Object, lw *cache.ListWatch) (cache.Store, cache.Controller) {
	var listWatch *cache.ListWatch
	if lw != nil {
		listWatch = lw
	} else {
		listWatch = cache.NewListWatchFromClient(w.KubeClient.CoreV1().RESTClient(), resource.String(), apiv1.NamespaceAll, fields.Everything())
	}

	return cache.NewInformer(
		listWatch,
		object,
		w.SyncPeriod,
		eventHandlerFuncs(w),
	)
}

func (w *Watcher) CacheStore(resource events.ObjectType, object runtime.Object, lw *cache.ListWatch) (cache.Store, cache.Controller) {
	if lw == nil {
		lw = cache.NewListWatchFromClient(w.KubeClient.CoreV1().RESTClient(), resource.String(), apiv1.NamespaceAll, fields.Everything())
	}

	return stash.NewInformerPopulated(
		lw,
		object,
		w.SyncPeriod,
		eventHandlerFuncs(w),
	)
}

func (w *Watcher) CacheIndexer(resource events.ObjectType, object runtime.Object, lw *cache.ListWatch, indexers cache.Indexers) (cache.Indexer, cache.Controller) {
	if lw == nil {
		lw = cache.NewListWatchFromClient(w.KubeClient.CoreV1().RESTClient(), resource.String(), apiv1.NamespaceAll, fields.Everything())
	}
	if indexers == nil {
		indexers = cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}
	}

	return stash.NewIndexerInformerPopulated(
		lw,
		object,
		w.SyncPeriod,
		eventHandlerFuncs(w),
		indexers,
	)
}

func eventHandlerFuncs(k *Watcher) cache.ResourceEventHandlerFuncs {
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			e := events.New(events.Added, obj)
			k.Dispatch(e)
		},
		DeleteFunc: func(obj interface{}) {
			e := events.New(events.Deleted, obj)
			k.Dispatch(e)
		},
		UpdateFunc: func(old, new interface{}) {
			if !reflect.DeepEqual(old, new) {
				e := events.New(events.Updated, old, new)
				k.Dispatch(e)
			}
		},
	}
}
