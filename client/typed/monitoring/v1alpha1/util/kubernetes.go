package util

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/appscode/kutil"
	api "github.com/appscode/searchlight/apis/monitoring/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func GetGroupVersionKind(v interface{}) schema.GroupVersionKind {
	return api.SchemeGroupVersion.WithKind(kutil.GetKind(v))
}

func AssignTypeKind(v interface{}) error {
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return fmt.Errorf("%v must be a pointer", v)
	}

	switch u := v.(type) {
	case *api.ClusterAlert:
		u.APIVersion = api.SchemeGroupVersion.String()
		u.Kind = kutil.GetKind(v)
		return nil
	case *api.NodeAlert:
		u.APIVersion = api.SchemeGroupVersion.String()
		u.Kind = kutil.GetKind(v)
		return nil
	case *api.PodAlert:
		u.APIVersion = api.SchemeGroupVersion.String()
		u.Kind = kutil.GetKind(v)
		return nil
	}
	return errors.New("unknown api object type")
}
