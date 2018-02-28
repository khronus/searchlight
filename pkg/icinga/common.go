package icinga

import (
	"encoding/json"
	"fmt"
	"strings"

	api "github.com/appscode/searchlight/apis/monitoring/v1alpha1"
	"github.com/pkg/errors"
)

type commonHost struct {
	IcingaClient *Client
}

func (h *commonHost) EnsureIcingaHost(kh IcingaHost) error {
	host, err := kh.Name()
	if err != nil {
		return errors.WithStack(err)
	}
	resp := h.IcingaClient.Objects().Hosts(host).Get([]string{}).Do()
	if resp.Status == 200 {
		return nil
	}
	obj := IcingaObject{
		Templates: []string{"generic-host"},
		Attrs: map[string]interface{}{
			"address": kh.IP,
		},
	}
	jsonStr, err := json.Marshal(obj)
	if err != nil {
		return errors.WithStack(err)
	}

	resp = h.IcingaClient.Objects().Hosts(host).Create([]string{}, string(jsonStr)).Do()
	if resp.Err != nil {
		return errors.WithStack(resp.Err)
	}
	if resp.Status != 200 {
		return errors.Errorf("Can't create Icinga host: %d", resp.Status)
	}
	return nil
}

func (h *commonHost) DeleteIcingaHost(kh IcingaHost) error {
	param := map[string]string{
		"cascade": "1",
	}
	host, err := kh.Name()
	if err != nil {
		return errors.WithStack(err)
	}

	in := fmt.Sprintf(`{"filter": "match(\"%s\",host.name)"}`, host)
	var respService ResponseObject
	if _, err := h.IcingaClient.Objects().Service("").Update([]string{}, in).Do().Into(&respService); err != nil {
		return errors.WithMessage(err, "Can't get Icinga service")
	}

	if len(respService.Results) <= 1 {
		resp := h.IcingaClient.Objects().Hosts("").Delete([]string{}, in).Params(param).Do()
		if resp.Err != nil {
			return errors.WithMessage(err, "Can't delete Icinga host")
		}
	}
	return nil
}

// createIcingaServiceForCluster
func (h *commonHost) CreateIcingaService(svc string, kh IcingaHost, attrs map[string]interface{}) error {
	obj := IcingaObject{
		Templates: []string{"generic-service"},
		Attrs:     attrs,
	}
	jsonStr, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	host, err := kh.Name()
	if err != nil {
		return errors.WithStack(err)
	}
	resp := h.IcingaClient.Objects().Service(host).Create([]string{svc}, string(jsonStr)).Do()
	if resp.Err != nil {
		return errors.WithStack(resp.Err)
	}
	if resp.Status == 200 {
		return nil
	}
	if strings.Contains(string(resp.ResponseBody), "already exists") {
		return nil
	}
	return errors.Errorf("Can't create Icinga service %d", resp.Status)
}

func (h *commonHost) UpdateIcingaService(svc string, kh IcingaHost, attrs map[string]interface{}) error {
	obj := IcingaObject{
		Templates: []string{"generic-service"},
		Attrs:     attrs,
	}
	jsonStr, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	host, err := kh.Name()
	if err != nil {
		return errors.WithStack(err)
	}
	resp := h.IcingaClient.Objects().Service(host).Update([]string{svc}, string(jsonStr)).Do()
	if resp.Err != nil {
		return errors.WithStack(resp.Err)
	}
	if resp.Status != 200 {
		return errors.Errorf("Can't update Icinga service; %d", resp.Status)
	}
	return nil
}

func (h *commonHost) DeleteIcingaService(svc string, kh IcingaHost) error {
	param := map[string]string{
		"cascade": "1",
	}
	in := h.IcingaServiceSearchQuery(svc, kh)
	resp := h.IcingaClient.Objects().Service("").Delete([]string{}, in).Params(param).Do()
	if resp.Err != nil {
		return errors.WithStack(resp.Err)
	}
	if resp.Status == 200 || resp.Status == 404 {
		return nil
	}
	return errors.Errorf("Fail to delete service: %d", resp.Status)
}

func (h *commonHost) CheckIcingaService(svc string, kh IcingaHost) (bool, error) {
	in := h.IcingaServiceSearchQuery(svc, kh)
	var respService ResponseObject

	if _, err := h.IcingaClient.Objects().Service("").Get([]string{}, in).Do().Into(&respService); err != nil {
		return true, errors.WithMessage(err, "Can't check icinga service")
	}
	return len(respService.Results) > 0, nil
}

func (h *commonHost) IcingaServiceSearchQuery(svc string, kids ...IcingaHost) string {
	matchHost := ""
	for i, kh := range kids {
		if i > 0 {
			matchHost = matchHost + "||"
		}
		host, _ := kh.Name()
		matchHost = matchHost + fmt.Sprintf(`match(\"%s\",host.name)`, host)
	}
	return fmt.Sprintf(`{"filter": "(%s)&&match(\"%s\",service.name)"}`, matchHost, svc)
}

func (h *commonHost) EnsureIcingaNotification(alert api.Alert, kh IcingaHost) error {
	obj := IcingaObject{
		Templates: []string{"icinga2-notifier-template"},
		Attrs: map[string]interface{}{
			"interval": int(alert.GetAlertInterval().Seconds()),
			"users":    []string{"searchlight_user"},
		},
	}
	jsonStr, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	host, err := kh.Name()
	if err != nil {
		return errors.WithStack(err)
	}

	var has bool
	var verb string
	var resp *APIResponse

	if !has {
		verb = "create"
		resp = h.IcingaClient.Objects().
			Notifications(host).
			Create([]string{alert.GetName(), alert.GetName()}, string(jsonStr)).
			Do()
	} else {
		verb = "update"
		resp = h.IcingaClient.Objects().
			Notifications(host).
			Update([]string{alert.GetName(), alert.GetName()}, string(jsonStr)).
			Do()
	}

	if resp.Err != nil {
		return errors.WithStack(resp.Err)
	}
	if resp.Status != 200 {
		return errors.Errorf("Can't %s Icinga notification: %d", verb, resp.Status)
	}

	return nil
}
