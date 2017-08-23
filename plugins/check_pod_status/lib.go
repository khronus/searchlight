package check_pod_status

import (
	"fmt"
	"os"

	"github.com/appscode/go/flags"
	"github.com/appscode/searchlight/pkg/icinga"
	"github.com/appscode/searchlight/pkg/util"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/client-go/pkg/api/v1"
)

type Request struct {
	Host string
}

type objectInfo struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Status    string `json:"status,omitempty"`
}

type serviceOutput struct {
	Objects []*objectInfo `json:"objects,omitempty"`
	Message string        `json:"message,omitempty"`
}

func CheckPodStatus(req *Request) (icinga.State, interface{}) {
	kubeClient, err := util.NewClient()
	if err != nil {
		return icinga.UNKNOWN, err
	}

	host, err := icinga.ParseHost(req.Host)
	if err != nil {
		fmt.Fprintln(os.Stdout, icinga.WARNING, "Invalid icinga host.name")
		os.Exit(3)
	}
	if host.Type != icinga.TypePod {
		fmt.Fprintln(os.Stdout, icinga.WARNING, "Invalid icinga host type")
		os.Exit(3)
	}

	pod, err := kubeClient.Client.CoreV1().Pods(host.AlertNamespace).Get(host.ObjectName, metav1.GetOptions{})
	if err != nil {
		return icinga.UNKNOWN, err
	}

	if ok, err := PodRunningAndReady(*pod); !ok {
		return icinga.CRITICAL, err
	}
	return icinga.OK, pod.Status.Phase
}

// PodRunningAndReady returns whether a pod is running and each container has
// passed it's ready state.
func PodRunningAndReady(pod apiv1.Pod) (bool, error) {
	switch pod.Status.Phase {
	case apiv1.PodFailed, apiv1.PodSucceeded:
		return false, fmt.Errorf("pod completed")
	case apiv1.PodRunning:
		for _, cond := range pod.Status.Conditions {
			if cond.Type != apiv1.PodReady {
				continue
			}
			return cond.Status == apiv1.ConditionTrue, nil
		}
		return false, fmt.Errorf("pod ready condition not found")
	}
	return false, nil
}

func NewCmd() *cobra.Command {
	var req Request
	c := &cobra.Command{
		Use:     "check_pod_status",
		Short:   "Check Kubernetes Pod(s) status",
		Example: "",

		Run: func(cmd *cobra.Command, args []string) {
			flags.EnsureRequiredFlags(cmd, "host")
			icinga.Output(CheckPodStatus(&req))
		},
	}
	c.Flags().StringVarP(&req.Host, "host", "H", "", "Icinga host name")
	return c
}
