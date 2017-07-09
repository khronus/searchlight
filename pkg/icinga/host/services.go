package host

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/appscode/errors"
	icinga "github.com/appscode/searchlight/pkg/icinga/client"
)

const (
	HostTypeLocalhost = "localhost"
	HostTypeNode      = "node"
	HostTypePod       = "pod"
)

// createIcingaServiceForCluster
func CreateIcingaService(icingaClient *icinga.IcingaClient, mp map[string]interface{}, object *KubeObjectInfo, serviceName string) error {
	var obj IcingaObject
	obj.Templates = []string{"generic-service"}
	obj.Attrs = mp
	jsonStr, err := json.Marshal(obj)
	if err != nil {
		return errors.New().WithCause(err).Err()
	}

	resp := icingaClient.Objects().Service(object.Name).Create([]string{serviceName}, string(jsonStr)).Do()
	if resp.Err != nil {
		return errors.New().WithCause(resp.Err).Err()
	}

	if resp.Status == 200 {
		return nil
	}
	if strings.Contains(string(resp.ResponseBody), "already exists") {
		return nil
	}

	return errors.New("Can't create Icinga service").Err()
}

func UpdateIcingaService(icingaClient *icinga.IcingaClient, mp map[string]interface{}, object *KubeObjectInfo, icignaService string) error {
	var obj IcingaObject
	obj.Templates = []string{"generic-service"}
	obj.Attrs = mp
	jsonStr, err := json.Marshal(obj)
	if err != nil {
		return errors.New().WithCause(err).Err()
	}
	resp := icingaClient.Objects().Service(object.Name).Update([]string{icignaService}, string(jsonStr)).Do()
	if resp.Err != nil {
		return errors.New().WithCause(resp.Err).Err()
	}

	if resp.Status != 200 {
		return errors.New("Can't update Icinga service").Err()
	}
	return nil
}

func DeleteIcingaService(icingaClient *icinga.IcingaClient, objectList []*KubeObjectInfo, icingaServiceName string) error {
	param := map[string]string{
		"cascade": "1",
	}
	in := IcingaServiceSearchQuery(icingaServiceName, objectList)
	resp := icingaClient.Objects().Service("").Delete([]string{}, in).Params(param).Do()

	if resp.Err != nil {
		return errors.New().WithCause(resp.Err).Err()
	}
	if resp.Status == 200 {
		return nil
	}
	return errors.New("Fail to delete service").Err()
}

func CheckIcingaService(icingaClient *icinga.IcingaClient, icingaServiceName string, objectList []*KubeObjectInfo) (bool, error) {
	in := IcingaServiceSearchQuery(icingaServiceName, objectList)
	var respService ResponseObject

	if _, err := icingaClient.Objects().Service("").Get([]string{}, in).Do().Into(&respService); err != nil {
		return true, errors.New("can't check icinga service").Err()
	}
	return len(respService.Results) > 0, nil
}

func IcingaServiceSearchQuery(icingaServiceName string, objectList []*KubeObjectInfo) string {
	matchHost := ""
	for id, object := range objectList {
		if id > 0 {
			matchHost = matchHost + "||"
		}
		matchHost = matchHost + fmt.Sprintf(`match(\"%s\",host.name)`, object.Name)
	}
	return fmt.Sprintf(`{"filter": "(%s)&&match(\"%s\",service.name)"}`, matchHost, icingaServiceName)
}
