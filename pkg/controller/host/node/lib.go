package node

import (
	"fmt"
	"regexp"

	"github.com/appscode/errors"
	aci "github.com/appscode/searchlight/api"
	"github.com/appscode/searchlight/data"
	"github.com/appscode/searchlight/pkg/controller/host"
	"github.com/appscode/searchlight/pkg/controller/host/extpoints"
	"github.com/appscode/searchlight/pkg/controller/types"
)

func init() {
	extpoints.IcingaHostTypes.Register(new(icingaHost), host.HostTypeNode)
}

type icingaHost struct {
}

type biblio struct {
	*types.Option
}

func (p *icingaHost) CreateAlert(ctx *types.Option, specificObject string) error {
	return (&biblio{ctx}).create(specificObject)
}

func (p *icingaHost) UpdateAlert(ctx *types.Option) error {
	return (&biblio{ctx}).update()
}

func (p *icingaHost) DeleteAlert(ctx *types.Option, specificObject string) error {
	return (&biblio{ctx}).delete(specificObject)
}

//-----------------------------------------------------
// set Alert in Icinga LocalHost
func (b *biblio) create(specificObject string) error {
	alertSpec := b.Resource.Spec

	if alertSpec.Check == "" {
		return errors.New("Invalid request").Err()
	}

	// Get Icinga Host Info
	objectList, err := host.GetObjectList(b.KubeClient, alertSpec.Check, host.HostTypeNode, b.Resource.Namespace, b.ObjectType, b.ObjectName, specificObject)
	if err != nil {
		return errors.New().WithCause(err).Err()
	}

	var has bool
	if has, err = host.CheckIcingaService(b.IcingaClient, b.Resource.Name, objectList); err != nil {
		return errors.New().WithCause(err).Err()
	}
	if has {
		return nil
	}

	// Create Icinga Host
	if err := host.CreateIcingaHost(b.IcingaClient, objectList, b.Resource.Namespace); err != nil {
		return errors.New().WithCause(err).Err()
	}

	if err := b.createIcingaService(objectList); err != nil {
		return errors.New().WithCause(err).Err()
	}

	if err := host.CreateIcingaNotification(b.IcingaClient, b.Resource, objectList); err != nil {
		return errors.New().WithCause(err).Err()
	}

	return nil
}

func setParameterizedVariables(alertSpec aci.PodAlertSpec, objectName string, commandVars map[string]data.CommandVar, mp map[string]interface{}) (map[string]interface{}, error) {
	for key, val := range alertSpec.Vars {
		if v, found := commandVars[key]; found {
			if !v.Parameterized {
				continue
			}

			reg, err := regexp.Compile("nodename[ ]*=[ ]*'[?]'")
			if err != nil {
				return nil, errors.New().WithCause(err).Err()
			}
			mp[host.IVar(key)] = reg.ReplaceAllString(val.(string), fmt.Sprintf("nodename='%s'", objectName))
		} else {
			return nil, errors.Newf("variable %v not found", key).Err()
		}
	}
	return mp, nil
}

func (b *biblio) createIcingaService(objectList []*host.KubeObjectInfo) error {
	alertSpec := b.Resource.Spec

	mp := make(map[string]interface{})
	mp["check_command"] = alertSpec.Check
	if alertSpec.CheckInterval.Seconds() > 0 {
		mp["check_interval"] = alertSpec.CheckInterval.Seconds()
	}

	commandVars := b.IcingaData[alertSpec.Check].VarInfo
	for key, val := range alertSpec.Vars {
		if v, found := commandVars[key]; found {
			if v.Parameterized {
				continue
			}
			mp[host.IVar(key)] = val
		}
	}

	for _, object := range objectList {
		var err error
		if mp, err = setParameterizedVariables(alertSpec, object.Name, commandVars, mp); err != nil {
			return errors.New().WithCause(err).Err()
		}

		if err := host.CreateIcingaService(b.IcingaClient, mp, object, b.Resource.Name); err != nil {
			return errors.New().WithCause(err).Err()
		}
	}
	return nil
}

func (b *biblio) update() error {
	alertSpec := b.Resource.Spec

	// Get Icinga Host Info
	objectList, err := host.GetObjectList(b.KubeClient, alertSpec.Check, host.HostTypeNode, b.Resource.Namespace, b.ObjectType, b.ObjectName, "")
	if err != nil {
		return errors.New().WithCause(err).Err()
	}

	if err := b.updateIcingaService(objectList); err != nil {
		return errors.New().WithCause(err).Err()
	}

	if err := host.UpdateIcingaNotification(b.IcingaClient, b.Resource, objectList); err != nil {
		return errors.New().WithCause(err).Err()
	}
	return nil
}

func (b *biblio) updateIcingaService(objectList []*host.KubeObjectInfo) error {
	alertSpec := b.Resource.Spec

	mp := make(map[string]interface{})
	if alertSpec.CheckInterval.Seconds() > 0 {
		mp["check_interval"] = alertSpec.CheckInterval.Seconds()
	}

	commandVars := b.IcingaData[alertSpec.Check].VarInfo
	for key, val := range alertSpec.Vars {
		if v, found := commandVars[key]; found {
			if v.Parameterized {
				continue
			}
			mp[host.IVar(key)] = val
		}
	}

	for _, object := range objectList {
		var err error
		if mp, err = setParameterizedVariables(alertSpec, object.Name, commandVars, mp); err != nil {
			return errors.New().WithCause(err).Err()
		}

		if err := host.UpdateIcingaService(b.IcingaClient, mp, object, b.Resource.Name); err != nil {
			return errors.New().WithCause(err).Err()
		}
	}
	return nil
}

func (b *biblio) delete(specificObject string) error {
	alertSpec := b.Resource.Spec

	var objectList []*host.KubeObjectInfo
	if specificObject != "" {
		objectList = append(objectList, &host.KubeObjectInfo{Name: specificObject + "@" + b.Resource.Namespace})
	} else {
		// Get Icinga Host Info
		var err error
		objectList, err = host.GetObjectList(b.KubeClient, alertSpec.Check, host.HostTypeNode,
			b.Resource.Namespace, b.ObjectType, b.ObjectName, specificObject)
		if err != nil {
			return errors.New().WithCause(err).Err()
		}

	}

	if err := host.DeleteIcingaService(b.IcingaClient, objectList, b.Resource.Name); err != nil {
		return errors.New().WithCause(err).Err()
	}

	for _, object := range objectList {
		if err := host.DeleteIcingaHost(b.IcingaClient, object.Name); err != nil {
			return errors.New().WithCause(err).Err()
		}
	}
	return nil
}
