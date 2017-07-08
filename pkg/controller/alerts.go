package controller

import (
	"errors"
	"fmt"
	"reflect"

	acrt "github.com/appscode/go/runtime"
	"github.com/appscode/log"
	tapi "github.com/appscode/searchlight/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	apiv1 "k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/tools/cache"
)

// Blocks caller. Intended to be called as a Go routine.
func (c *Controller) WatchAlerts() {
	defer acrt.HandleCrash()

	lw := &cache.ListWatch{
		ListFunc: func(opts metav1.ListOptions) (runtime.Object, error) {
			return c.extClient.Alerts(apiv1.NamespaceAll).List(metav1.ListOptions{})
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return c.extClient.Alerts(apiv1.NamespaceAll).Watch(metav1.ListOptions{})
		},
	}
	_, ctrl := cache.NewInformer(lw,
		&tapi.Alert{},
		c.syncPeriod,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				if resource, ok := obj.(*tapi.Alert); ok {
					fmt.Println(resource.Name)
				}
			},
			UpdateFunc: func(old, new interface{}) {
				oldObj, ok := old.(*tapi.Alert)
				if !ok {
					log.Errorln(errors.New("Invalid Alert object"))
					return
				}
				newObj, ok := new.(*tapi.Alert)
				if !ok {
					log.Errorln(errors.New("Invalid Alert object"))
					return
				}
				if !reflect.DeepEqual(oldObj, newObj) {
				}
			},
			DeleteFunc: func(obj interface{}) {
				if resource, ok := obj.(*tapi.Alert); ok {
					fmt.Println(resource.Name)
				}
			},
		},
	)
	ctrl.Run(wait.NeverStop)
}