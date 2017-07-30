> New to Searchlight? Please start [here](/docs/tutorials/README.md).

# ClusterAlerts

## What is ClusterAlert
A `ClusterAlert` is a Kubernetes `Third Party Object` (TPR). It provides declarative configuration of [Icinga services](https://www.icinga.com/docs/icinga2/latest/doc/09-object-types/#service) for cluster level alerts in a Kubernetes native way. You only need to describe the desired check command and notifier in a ClusterAlert object, and the Searchlight operator will create Icinga2 hosts, services and notifications to the desired state for you.

## ClusterAlert Spec
As with all other Kubernetes objects, a ClusterAlert needs `apiVersion`, `kind`, and `metadata` fields. It also needs a `.spec` section. Below is an example ClusterAlert object.

```yaml
apiVersion: monitoring.appscode.com/v1alpha1
kind: ClusterAlert
metadata:
  name: pod-exists-demo-0
  namespace: demo
spec:
  check: pod_exists
  vars:
    selector: app=nginx
    count: 2
  checkInterval: 5m
  alertInterval: 3m
  notifierSecretName: notifier-config
  receivers:
  - notifier: mailgun
    state: WARNING
    to: ["ops@example.com"]
  - notifier: twilio
    state: CRITICAL
    to: ["+1-234-567-8901"]
```

This object will do the followings:

- This Alert is set on pods with matching label `app=nginx` in `demo` namespace.
- Check command `pod_volume` will be applied on volume named `webstore`.
- Icinga will check for volume size every 60s.
- Notifications will be sent every 5m if any problem is detected, until acknowledged.
- When the disk is 70% full, it will reach `WARNING` state and emails will be sent to _ops@example.com_ via Mailgun as notification.
- When the disk is 95% full, it will reach `CRITICAL` state and SMSes will be sent to _+1-234-567-8901_ via Twilio as notification.

Any ClusterAlert object has 2 main sections:

### Check Command
Check commands are used by Icinga to periodically test some condition. If the test return positive appropriate notifications are sent. The following check commands are supported for pods:
- [any_http](any_http.md) - To check any HTTP response.
- [ca_cert](ca_cert.md) - To check expiration of CA certificate used by Kubernetes api server.
- [component_status](component_status.md) - To check Kubernetes component status.
- [event](event.md) - To check Kubernetes Warning events.
- [json_path](json_path.md) - To check any HTTP response by parsing as JSON using [jq](https://stedolan.github.io/jq/).
- [node_exists](node_count.md) - To check existence of Kubernetes nodes.
- [pod_exists](pod_exists.md) - To check existence of Kubernetes pods.

Each check command has a name specified in `spec.check` field. Optionally each check command can take one or more parameters. These are specified in `spec.vars` field. To learn about the available parameters for each check command, please visit their documentation. `spec.checkInterval` specifies how frequently Icinga will perform this check. Some examples are: 30s, 5m, 6h, etc.

### Notifiers
When a check fails, Icinga will keep sending notifications until acknowledged via IcingaWeb dashboard. `spec.alertInterval` specifies how frequently notifications are sent. Icinga can send notifications to different targets based on alert state. `spec.receivers` contains that list of targets:

| Name                       | Description                                                  |
|----------------------------|--------------------------------------------------------------|
| `spec.receivers[*].state`  | `Required` Name of state for which notification will be sent |
| `spec.receivers[*].to`     | `Required` To whom notifications will be sent                |
| `spec.receivers[*].method` | `Required` How this notification will be sent                |


## Icinga Objects
You can skip this section if you are unfamiliar with how Icinga works. Searchlight operator watches for ClusterAlert objects and turns them into [Icinga objects](https://www.icinga.com/docs/icinga2/latest/doc/09-object-types/) accordingly. A single [Icinga Host](https://www.icinga.com/docs/icinga2/latest/doc/09-object-types/#host) is created with the name `{namespace}@cluster` and address `127.0.0.1` for all ClusterAlerts in a Kubernetes namespace. Now for each ClusterAlert, an [Icinga service](https://www.icinga.com/docs/icinga2/latest/doc/09-object-types/#service) is created with name matching the ClusterAlert name.
