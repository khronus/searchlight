package notifier

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	api "github.com/appscode/searchlight/apis/monitoring/v1alpha1"
	"github.com/appscode/searchlight/pkg/icinga"
)

func RenderSubject(alert api.Alert, req *Request) string {
	if strings.ToUpper(req.Type) == EventTypeAcknowledgement {
		return fmt.Sprintf("Problem Acknowledged: Service [%s] for [%s] is in \"%s\" state", alert.GetName(), req.HostName, req.State)
	} else if strings.ToUpper(req.Type) == EventTypeRecovery {
		return fmt.Sprintf("Problem Recovered: Service [%s] for [%s] is in \"%s\" state.", alert.GetName(), req.HostName, req.State)
	} else if strings.ToUpper(req.Type) == EventTypeProblem {
		return fmt.Sprintf("Problem Detected: Service [%s] for [%s] is in \"%s\" state.", alert.GetName(), req.HostName, req.State)
	} else {
		return fmt.Sprintf("Service [%s] for [%s] is in \"%s\" state.", alert.GetName(), req.HostName, req.State)
	}
}

type TemplateData struct {
	AlertNamespace     string
	AlertType          string
	AlertName          string
	ObjectName         string
	IcingaHostName     string
	IcingaServiceName  string
	IcingaCheckCommand string
	IcingaType         string
	IcingaState        string
	IcingaOutput       string
	Author             string
	Comment            string
	IcingaTime         time.Time
}

func RenderMail(alert api.Alert, req *Request) (string, error) {
	host, err := icinga.ParseHost(req.HostName)
	if err != nil {
		return "", err
	}
	data := TemplateData{
		AlertName:          alert.GetName(),
		AlertNamespace:     host.AlertNamespace,
		AlertType:          host.Type,
		ObjectName:         host.ObjectName,
		IcingaHostName:     req.HostName,
		IcingaServiceName:  alert.GetName(),
		IcingaCheckCommand: alert.Command(),
		IcingaType:         req.Type,
		IcingaState:        strings.ToUpper(req.State),
		IcingaOutput:       req.Output,
		Author:             req.Author,
		Comment:            req.Comment,
		IcingaTime:         req.Time,
	}

	var buf bytes.Buffer
	if err = mailTemplate.Execute(&buf, data); err != nil {
		return "", err
	}
	config := buf.String()
	return config, nil
}
