package controller

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/appscode/errors"
	"github.com/appscode/log"
	"github.com/appscode/searchlight/pkg/client/icinga"
	"github.com/appscode/searchlight/pkg/controller/types"
	apiv1 "k8s.io/client-go/pkg/api/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (b *IcingaController) Acknowledge(event *apiv1.Event) error {
	icingaService := b.ctx.Resource.Name

	var message types.AlertEventMessage
	err := json.Unmarshal([]byte(event.Message), &message)
	if err != nil {
		return errors.New().WithCause(err).Err()
	}

	if event.Source.Host == "" {
		return errors.New("Icinga hostname missing").Err()
	}
	if err = acknowledgeIcingaNotification(b.ctx.IcingaClient, event.Source.Host, icingaService, message.Comment, message.UserName); err != nil {
		return errors.New().WithCause(err).Err()
	}

	if event.Annotations == nil {
		event.Annotations = make(map[string]string)
	}

	timestamp := v1.NewTime(time.Now().UTC())
	event.Annotations[types.AcknowledgeTimestamp] = timestamp.String()

	if _, err = b.ctx.KubeClient.Core().Events(event.Namespace).Update(event); err != nil {
		return errors.New().WithCause(err).Err()
	}
	return nil
}

func acknowledgeIcingaNotification(client *icinga.IcingaClient, icingaHostName, icingaServiceName, comment, username string) error {
	mp := make(map[string]interface{})
	mp["type"] = "Service"
	mp["filter"] = fmt.Sprintf(`service.name == "%s" && host.name == "%s"`, icingaServiceName, icingaHostName)
	mp["comment"] = comment
	mp["notify"] = true
	mp["author"] = username

	jsonStr, err := json.Marshal(mp)
	if err != nil {
		return errors.New().WithCause(err).Err()
	}
	resp := client.Actions("acknowledge-problem").Update([]string{}, string(jsonStr)).Do()
	if resp.Status == 200 {
		log.Debugln("[Icinga] Problem acknowledged")
		return nil
	}
	return errors.New("[Icinga] Problem acknowledged Error").Err()
}
