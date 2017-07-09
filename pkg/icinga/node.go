package icinga

import (
	"fmt"
	"regexp"

	"github.com/appscode/errors"
	tapi "github.com/appscode/searchlight/api"
	tcs "github.com/appscode/searchlight/client/clientset"
	"github.com/appscode/searchlight/pkg/icinga"
	clientset "k8s.io/client-go/kubernetes"
	apiv1 "k8s.io/client-go/pkg/api/v1"
)

type NodeHost struct {
	commonHost

	KubeClient clientset.Interface
	ExtClient  tcs.ExtensionInterface
	//*types.Context
}

func NewNodeHost(kubeClient clientset.Interface, extClient tcs.ExtensionInterface, IcingaClient *icinga.Client) *NodeHost {
	return &NodeHost{
		KubeClient: kubeClient,
		ExtClient:  extClient,
		commonHost: commonHost{
			IcingaClient: IcingaClient,
		},
	}
}

func (h *NodeHost) GetObject(alert tapi.NodeAlert, node apiv1.Node) KHost {
	nodeIP := "127.0.0.1"
	for _, ip := range node.Status.Addresses {
		if ip.Type == internalIP {
			nodeIP = ip.Address
			break
		}
	}
	return KHost{Name: node.Name + "@" + alert.Namespace, IP: nodeIP}
}

func (h *NodeHost) expandVars(alertSpec tapi.NodeAlertSpec, kh KHost, attrs map[string]interface{}) error {
	commandVars := tapi.NodeCommands[alertSpec.Check].Vars
	for key, val := range alertSpec.Vars {
		if v, found := commandVars[key]; found {
			if v.Parameterized {
				reg, err := regexp.Compile("nodename[ ]*=[ ]*'[?]'")
				if err != nil {
					return err
				}
				attrs[IVar(key)] = reg.ReplaceAllString(val.(string), fmt.Sprintf("nodename='%s'", kh.Name))
			} else {
				attrs[IVar(key)] = val
			}
		} else {
			return errors.Newf("variable %v not found", key).Err()
		}
	}
	return nil
}

// set Alert in Icinga LocalHost
func (h *NodeHost) Create(alert tapi.NodeAlert, node apiv1.Node) error {
	alertSpec := alert.Spec
	kh := h.GetObject(alert, node)

	if has, err := h.CheckIcingaService(alert.Name, kh); err != nil || has {
		return err
	}

	if err := h.CreateIcingaHost(kh); err != nil {
		return errors.FromErr(err).Err()
	}

	attrs := make(map[string]interface{})
	attrs["check_command"] = alertSpec.Check
	if alertSpec.CheckInterval.Seconds() > 0 {
		attrs["check_interval"] = alertSpec.CheckInterval.Seconds()
	}
	if err := h.expandVars(alertSpec, kh, attrs); err != nil {
		return err
	}
	if err := h.CreateIcingaService(alert.Name, kh, attrs); err != nil {
		return errors.FromErr(err).Err()
	}

	return h.CreateIcingaNotification(alert, kh)
}

func (h *NodeHost) Update(alert tapi.NodeAlert, node apiv1.Node) error {
	alertSpec := alert.Spec
	kh := h.GetObject(alert, node)

	attrs := make(map[string]interface{})
	if alertSpec.CheckInterval.Seconds() > 0 {
		attrs["check_interval"] = alertSpec.CheckInterval.Seconds()
	}
	if err := h.expandVars(alertSpec, kh, attrs); err != nil {
		return err
	}
	if err := h.UpdateIcingaService(alert.Name, kh, attrs); err != nil {
		return errors.FromErr(err).Err()
	}

	return h.UpdateIcingaNotification(alert, kh)
}

func (h *NodeHost) Delete(alert tapi.NodeAlert, node apiv1.Node) error {
	kh := h.GetObject(alert, node)

	if err := h.DeleteIcingaService(alert.Name, kh); err != nil {
		return errors.FromErr(err).Err()
	}
	return h.DeleteIcingaHost(kh.Name)
}
