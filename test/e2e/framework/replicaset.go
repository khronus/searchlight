package framework

import (
	"github.com/appscode/go/crypto/rand"
	"github.com/appscode/go/types"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/client-go/pkg/api/v1"
	extensions "k8s.io/client-go/pkg/apis/extensions/v1beta1"
)

func (f *Invocation) ReplicaSet() *extensions.ReplicaSet {
	return &extensions.ReplicaSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      rand.WithUniqSuffix("replicaset"),
			Namespace: f.namespace,
			Labels: map[string]string{
				"app": f.app,
			},
		},
		Spec: extensions.ReplicaSetSpec{
			Replicas: types.Int32P(1),
			Template: f.PodTemplate(),
		},
	}
}

func (f *Framework) CreateReplicaSet(obj *extensions.ReplicaSet) (*extensions.ReplicaSet, error) {
	return f.kubeClient.ExtensionsV1beta1().ReplicaSets(obj.Namespace).Create(obj)
}

func (f *Framework) DeleteReplicaSet(meta metav1.ObjectMeta) error {
	return f.kubeClient.ExtensionsV1beta1().ReplicaSets(meta.Namespace).Delete(meta.Name, deleteInForeground())
}

func (f *Framework) EventuallyReplicaSetRunning(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(func() *apiv1.PodList {
		obj, err := f.kubeClient.ExtensionsV1beta1().ReplicaSets(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())
		podList, err := f.GetPodList(obj)
		Expect(err).NotTo(HaveOccurred())
		return podList
	})
}
