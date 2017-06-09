package check_kube_event

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/appscode/go/flags"
	"github.com/appscode/searchlight/pkg/client/k8s"
	"github.com/appscode/searchlight/util"
	"github.com/spf13/cobra"
	apiv1 "k8s.io/client-go/pkg/api/v1"
	"k8s.io/kubernetes/pkg/fields"
)

type eventInfo struct {
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Kind      string `json:"kind,omitempty"`
	Count     int32  `json:"count,omitempty"`
	Reason    string `json:"reason,omitempty"`
	Message   string `json:"message,omitempty"`
}

type serviceOutput struct {
	Events  []*eventInfo `json:"events,omitempty"`
	Message string       `json:"message,omitempty"`
}

func CheckKubeEvent(req *Request) (util.IcingaState, interface{}) {
	kubeClient, err := k8s.NewClient()
	if err != nil {
		return util.Unknown, err
	}

	eventInfoList := make([]*eventInfo, 0)
	field := fields.OneTermEqualSelector(apiv1.EventTypeField, apiv1.EventTypeWarning)

	checkTime := time.Now().Add(-(req.CheckInterval + req.ClockSkew))

	eventList, err := kubeClient.Client.Core().Events(apiv1.NamespaceAll).List(
		apiv1.ListOptions{
			FieldSelector: field,
		},
	)
	if err != nil {
		return util.Unknown, err
	}

	for _, event := range eventList.Items {
		if checkTime.Before(event.LastTimestamp.Time) {
			eventInfoList = append(eventInfoList,
				&eventInfo{
					Name:      event.InvolvedObject.Name,
					Namespace: event.InvolvedObject.Namespace,
					Kind:      event.InvolvedObject.Kind,
					Count:     event.Count,
					Reason:    event.Reason,
					Message:   event.Message,
				},
			)
		}
	}

	if len(eventInfoList) == 0 {
		return util.Ok, "All events look Normal"
	} else {
		output := &serviceOutput{
			Events:  eventInfoList,
			Message: fmt.Sprintf("Found %d Warning event(s)", len(eventInfoList)),
		}
		outputByte, err := json.MarshalIndent(output, "", "  ")
		if err != nil {
			return util.Unknown, err
		}
		return util.Warning, outputByte
	}
}

type Request struct {
	CheckInterval time.Duration
	ClockSkew     time.Duration
}

func NewCmd() *cobra.Command {
	var req Request

	c := &cobra.Command{
		Use:     "check_kube_event",
		Short:   "Check kubernetes events for all namespaces",
		Example: "",

		Run: func(cmd *cobra.Command, args []string) {
			flags.EnsureRequiredFlags(cmd, "check_interval")
			util.Output(CheckKubeEvent(&req))
		},
	}

	c.Flags().DurationVarP(&req.CheckInterval, "check_interval", "c", time.Second*0, "Icinga check_interval in duration. [Format: 30s, 5m]")
	c.Flags().DurationVarP(&req.ClockSkew, "clock_skew", "s", time.Second*30, "Add skew with check_interval in duration. [Default: 30s]")
	return c
}
