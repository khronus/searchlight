package framework

import (
	"github.com/appscode/go/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	apiv1 "k8s.io/client-go/pkg/api/v1"
	apps "k8s.io/client-go/pkg/apis/apps/v1beta1"
	extensions "k8s.io/client-go/pkg/apis/extensions/v1beta1"
)

func (f *Invocation) DeploymentAppSearchlight() *apps.Deployment {
	return &apps.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      f.name,
			Namespace: f.namespace,
			Labels: map[string]string{
				"app": "searchlight",
			},
		},
		Spec: apps.DeploymentSpec{
			Replicas: types.Int32P(1),
			Template: f.getSearchlightPodTemplate(),
		},
	}
}

func (f *Invocation) DeploymentExtensionSearchlight() *extensions.Deployment {
	return &extensions.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      f.name,
			Namespace: f.namespace,
			Labels: map[string]string{
				"app": "searchlight",
			},
		},
		Spec: extensions.DeploymentSpec{
			Replicas: types.Int32P(1),
			Template: f.getSearchlightPodTemplate(),
		},
	}
}

func (f *Invocation) ServiceSearchlight() *apiv1.Service {
	return &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      f.name,
			Namespace: f.namespace,
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{
				"app": "searchlight",
			},
			Type: apiv1.ServiceTypeLoadBalancer,
			Ports: []apiv1.ServicePort{
				{
					Name:       "api",
					Port:       5665,
					TargetPort: intstr.Parse("api"),
				},
				{
					Name:       "web",
					Port:       80,
					TargetPort: intstr.Parse("web"),
				},
			},
		},
	}
}

func (f *Invocation) getSearchlightPodTemplate() apiv1.PodTemplateSpec {
	return apiv1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app": "searchlight",
			},
		},
		Spec: apiv1.PodSpec{
			Containers: []apiv1.Container{
				{
					Name:            "icinga",
					Image:           "aerokite/icinga:e2e-test-k8s",
					ImagePullPolicy: apiv1.PullAlways,
					Ports: []apiv1.ContainerPort{
						{
							ContainerPort: 5665,
							Name:          "api",
						},
						{
							ContainerPort: 60006,
							Name:          "web",
						},
					},
					VolumeMounts: []apiv1.VolumeMount{
						{
							Name:      "data-volume",
							MountPath: "/var/pv",
						},
						{
							Name:      "script-volume",
							MountPath: "/var/db-script",
						},
						{
							Name:      "icingaconfig",
							MountPath: "/srv",
						},
					},
				},
				{
					Name:            "ido",
					Image:           "appscode/postgres:9.5-v3-db",
					ImagePullPolicy: apiv1.PullIfNotPresent,
					Args: []string{
						"basic",
						"./setup-db.sh",
					},
					Ports: []apiv1.ContainerPort{
						{
							ContainerPort: 5432,
							Name:          "ido",
						},
					},
					VolumeMounts: []apiv1.VolumeMount{
						{
							Name:      "data-volume",
							MountPath: "/var/pv",
						},
						{
							Name:      "script-volume",
							MountPath: "/var/db-script",
						},
					},
				},
			},
			Volumes: []apiv1.Volume{
				{
					Name: "data-volume",
					VolumeSource: apiv1.VolumeSource{
						EmptyDir: &apiv1.EmptyDirVolumeSource{},
					},
				},
				{
					Name: "script-volume",
					VolumeSource: apiv1.VolumeSource{
						EmptyDir: &apiv1.EmptyDirVolumeSource{},
					},
				},
				{
					Name: "icingaconfig",
					VolumeSource: apiv1.VolumeSource{
						GitRepo: &apiv1.GitRepoVolumeSource{
							Repository: "https://github.com/appscode/icinga-testconfig.git",
							Directory:  ".",
						},
					},
				},
			},
		},
	}
}
