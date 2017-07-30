> New to Searchlight? Please start [here](/docs/tutorials/README.md).

# Check node_status

Check command `node_status` is used to check status of Kubernetes Nodes.

## Spec
`env` check command has no variables. Execution of this command can result in following states:
- OK
- CRITICAL
- UNKNOWN


## Tutorial

### Before You Begin
At first, you need to have a Kubernetes cluster, and the kubectl command-line tool must be configured to communicate with your cluster. If you do not already have a cluster, you can create one by using [Minikube](https://github.com/kubernetes/minikube).

To keep things isolated, this tutorial uses a separate namespace called `demo` throughout this tutorial. Run the following command to prepare your cluster for this tutorial:

```console
$ kubectl create namespace demo
namespace "demo" created

~ $ kubectl get namespaces
NAME          STATUS    AGE
default       Active    6h
kube-public   Active    6h
kube-system   Active    6h
demo          Active    4m
```

### Create Alert
In this tutorial, we are going to create an alert to check `env`.
```yaml
$ cat ./docs/examples/cluster-alerts/env/demo-0.yaml

apiVersion: monitoring.appscode.com/v1alpha1
kind: NodeAlert
metadata:
  name: env-demo-0
  namespace: demo
spec:
  check: env
  checkInterval: 30s
  alertInterval: 2m
  notifierSecretName: notifier-config
  receivers:
  - notifier: mailgun
    state: CRITICAL
    to: ["ops@example.com"]
```
```console
$ kubectl apply -f ./docs/examples/cluster-alerts/env/demo-0.yaml
clusteralert "env-demo-0" created

$ kubectl describe clusteralert env-demo-0 -n demo
Name:		env-demo-0
Namespace:	demo
Labels:		<none>
Events:
  FirstSeen	LastSeen	Count	From			SubObjectPath	Type		Reason		Message
  ---------	--------	-----	----			-------------	--------	------		-------
  6m		6m		1	Searchlight operator			Warning		BadNotifier	Bad notifier config for NodeAlert: "env-demo-0". Reason: secrets "notifier-config" not found
  6m		6m		1	Searchlight operator			Normal		SuccessfulSync	Applied NodeAlert: "env-demo-0"
```

Voila! `env` command has been synced to Icinga2. Searchlight also logged a warning event, we have not created the notifier secret `notifier-config`. Please visit [here](/docs/tutorials/notifiers.md) to learn how to configure notifier secret. Now, open IcingaWeb2 in your browser. You should see a Icinga host `demo@cluster` and Icinga service `env-demo-0`.

![Demo of check_env](/docs/images/cluster-alerts/env/demo-0.gif)

### Cleaning up
To cleanup the Kubernetes resources created by this tutorial, run:
```console
$ kubectl delete ns demo
```

If you would like to uninstall Searchlight operator, please follow the steps [here](/docs/uninstall.md).


## Next Steps


#### Supported Kubernetes Objects

| Kubernetes Object | Icinga2 Host Type |
| :---:             | :---:             |
| cluster           | node              |
| nodes             | node              |

#### Supported Icinga2 State

#### Example
###### Command
```console
hyperalert check_node_status --host=ip-172-20-0-9.ec2.internal@default
# --host is provided by Icinga2
```
###### Output
```
OK: Node is Ready
```

##### Configure Alert Object
```yaml
# This alert will be set to all nodes individually
apiVersion: monitoring.appscode.com/v1alpha1
kind: NodeAlert
metadata:
  name: check-node-status
  namespace: demo
spec:
  check: node_status
  alertInterval: 2m
  checkInterval: 1m
  receivers:
  - notifier: mailgun
    state: CRITICAL
    to: ["ops@example.com"]

# To set alert on specific node, set following labels
# labels:
#   alert.appscode.com/objectType: nodes
#   alert.appscode.com/objectName: ip-172-20-0-9.ec2.internal
```

```yaml
$ cat ./docs/examples/node-alerts/node_status/demo-0.yaml

apiVersion: monitoring.appscode.com/v1alpha1
kind: NodeAlert
metadata:
  name: node-status-demo-0
  namespace: demo
spec:
  check: node_status
  checkInterval: 30s
  alertInterval: 2m
  notifierSecretName: notifier-config
  receivers:
  - notifier: mailgun
    state: CRITICAL
    to: ["ops@example.com"]
```
```console
$ kubectl apply -f ./docs/examples/node-alerts/node_status/demo-0.yaml
nodealert "node-status-demo-0" created

$ kubectl describe nodealert -n demo node-status-demo-0
Name:		node-status-demo-0
Namespace:	demo
Labels:		<none>
Events:
  FirstSeen	LastSeen	Count	From			SubObjectPath	Type		Reason		Message
  ---------	--------	-----	----			-------------	--------	------		-------
  6s		6s		1	Searchlight operator			Warning		BadNotifier	Bad notifier config for NodeAlert: "node-status-demo-0". Reason: secrets "notifier-config" not found
  6s		6s		1	Searchlight operator			Normal		SuccessfulSync	Applied NodeAlert: "node-status-demo-0"
```


```console
$ kubectl apply -f ./docs/examples/node-alerts/node_status/demo-1.yaml
nodealert "node-status-demo-1" created

$ kubectl get nodealert -n demo
NAME                 KIND
node-status-demo-1   NodeAlert.v1alpha1.monitoring.appscode.com

$ kubectl describe nodealert -n demo node-status-demo-1
Name:		node-status-demo-1
Namespace:	demo
Labels:		<none>
Events:
  FirstSeen	LastSeen	Count	From			SubObjectPath	Type		Reason		Message
  ---------	--------	-----	----			-------------	--------	------		-------
  33s		33s		1	Searchlight operator			Warning		BadNotifier	Bad notifier config for NodeAlert: "node-status-demo-1". Reason: secrets "notifier-config" not found
  33s		33s		1	Searchlight operator			Normal		SuccessfulSync	Applied NodeAlert: "node-status-demo-1"
```

```console
$ kubectl apply -f ./docs/examples/node-alerts/node_status/demo-2.yaml
nodealert "node-status-demo-2" created

$ kubectl describe nodealert -n demo node-status-demo-2
Name:		node-status-demo-2
Namespace:	demo
Labels:		<none>
Events:
  FirstSeen	LastSeen	Count	From			SubObjectPath	Type		Reason		Message
  ---------	--------	-----	----			-------------	--------	------		-------
  22s		22s		1	Searchlight operator			Warning		BadNotifier	Bad notifier config for NodeAlert: "node-status-demo-2". Reason: secrets "notifier-config" not found
  22s		22s		1	Searchlight operator			Normal		SuccessfulSync	Applied NodeAlert: "node-status-demo-2"
```
