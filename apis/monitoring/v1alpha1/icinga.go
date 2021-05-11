/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"context"
	"strings"
	"sync"

	"gomodules.xyz/notify/unified"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	CheckPodStatus = "pod-status"
	CheckPodVolume = "pod-volume"
	CheckPodExec   = "pod-exec"
)

const (
	CheckNodeVolume = "node-volume"
	CheckNodeStatus = "node-status"
)

const (
	CheckComponentStatus = "component-status"
	CheckJsonPath        = "json-path"
	CheckNodeExists      = "node-exists"
	CheckPodExists       = "pod-exists"
	CheckEvent           = "event"
	CheckCACert          = "ca-cert"
)

// +k8s:deepcopy-gen=false
type Registry struct {
	reg map[string]IcingaCommand
	mu  sync.RWMutex
}

func (c *Registry) Get(cmd string) (IcingaCommand, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.reg[cmd]
	return v, ok
}

func (c *Registry) Insert(cmd string, v IcingaCommand) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.reg[cmd] = v
}

func (c *Registry) Delete(cmd string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.reg, cmd)
}

// +k8s:deepcopy-gen=false
type IcingaCommand struct {
	Name   string      `protobuf:"bytes,1,opt,name=name"`
	Vars   *PluginVars `protobuf:"bytes,2,opt,name=vars"`
	States []string    `protobuf:"bytes,3,rep,name=states"`
}

var (
	PodCommands     = &Registry{reg: map[string]IcingaCommand{}}
	NodeCommands    = &Registry{reg: map[string]IcingaCommand{}}
	ClusterCommands = &Registry{reg: map[string]IcingaCommand{}}
)

func checkNotifiers(kc kubernetes.Interface, alert Alert) error {
	if alert.GetNotifierSecretName() == "" && len(alert.GetReceivers()) == 0 {
		return nil
	}
	secret, err := kc.CoreV1().Secrets(alert.GetNamespace()).Get(context.TODO(), alert.GetNotifierSecretName(), metav1.GetOptions{})
	if err != nil {
		return err
	}
	for _, r := range alert.GetReceivers() {
		_, err = unified.LoadVia(r.Notifier, func(key string) (value string, found bool) {
			var bytes []byte
			bytes, found = secret.Data[key]
			value = string(bytes)
			return
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func AlertType(t string) IncidentNotificationType {
	switch strings.ToUpper(t) {
	case "PROBLEM":
		return NotificationProblem
	case "ACKNOWLEDGEMENT":
		return NotificationAcknowledgement
	case "RECOVERY":
		return NotificationRecovery
	default:
		return NotificationCustom
	}
}
