package icinga

import (
	"bytes"
	"text/template"

	api "github.com/appscode/searchlight/apis/monitoring/v1alpha1"
	"github.com/pkg/errors"
	core "k8s.io/api/core/v1"
)

type NodeHost struct {
	commonHost

	//*types.Context
}

func NewNodeHost(IcingaClient *Client) *NodeHost {
	return &NodeHost{
		commonHost: commonHost{
			IcingaClient: IcingaClient,
		},
	}
}

func (h *NodeHost) getHost(namespace string, node *core.Node) IcingaHost {
	nodeIP := "127.0.0.1"
	for _, ip := range node.Status.Addresses {
		if ip.Type == internalIP {
			nodeIP = ip.Address
			break
		}
	}
	return IcingaHost{
		ObjectName:     node.Name,
		Type:           TypeNode,
		AlertNamespace: namespace,
		IP:             nodeIP,
	}
}

func (h *NodeHost) expandVars(alertSpec api.NodeAlertSpec, kh IcingaHost, attrs map[string]interface{}) error {
	commandVars := api.NodeCommands[alertSpec.Check].Vars
	for key, val := range alertSpec.Vars {
		if v, found := commandVars[key]; found {
			if v.Parameterized {
				type Data struct {
					NodeName string
					NodeIP   string
				}
				tmpl, err := template.New("").Parse(val)
				if err != nil {
					return err
				}
				var buf bytes.Buffer
				err = tmpl.Execute(&buf, Data{NodeName: kh.ObjectName, NodeIP: kh.IP})
				if err != nil {
					return err
				}
				attrs[IVar(key)] = buf.String()
			} else {
				attrs[IVar(key)] = val
			}
		} else {
			return errors.Errorf("variable %v not found", key)
		}
	}
	return nil
}

// set Alert in Icinga LocalHost
func (h *NodeHost) Reconcile(alert *api.NodeAlert, node *core.Node) error {
	alertSpec := alert.Spec
	kh := h.getHost(alert.Namespace, node)

	if err := h.EnsureIcingaHost(kh); err != nil {
		return err
	}

	has, err := h.CheckIcingaService(alert.Name, kh)
	if err != nil {
		return err
	}

	attrs := make(map[string]interface{})
	if alertSpec.CheckInterval.Seconds() > 0 {
		attrs["check_interval"] = alertSpec.CheckInterval.Seconds()
	}
	if err := h.expandVars(alertSpec, kh, attrs); err != nil {
		return err
	}

	if !has {
		attrs["check_command"] = alertSpec.Check
		if err := h.CreateIcingaService(alert.Name, kh, attrs); err != nil {
			return err
		}
	} else {
		if err := h.UpdateIcingaService(alert.Name, kh, attrs); err != nil {
			return err
		}
	}

	return h.ReconcileIcingaNotification(alert, kh)
}

func (h *NodeHost) Delete(alertNamespace, alertName string, node *core.Node) error {
	kh := h.getHost(alertNamespace, node)

	if err := h.DeleteIcingaService(alertName, kh); err != nil {
		return err
	}
	return h.DeleteIcingaHost(kh)
}
