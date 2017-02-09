package e2e

import (
	"fmt"
	"testing"

	"github.com/appscode/searchlight/pkg/controller/host"
	"github.com/appscode/searchlight/test/mini"
	"github.com/appscode/searchlight/test/util"
	"github.com/stretchr/testify/assert"
)

func TestMultipleAlerts(t *testing.T) {

	// Run KubeD
	// runKubeD(setIcingaClient bool)
	// Pass true to set IcingaClient in watcher
	watcher, err := runKubeD(true)
	if !assert.Nil(t, err) {
		return
	}
	fmt.Println("--> Running kubeD")

	// Create ReplicaSet
	fmt.Println("--> Creating ReplicaSet")
	replicaSet, err := mini.CreateReplicaSet(watcher, "default")
	if !assert.Nil(t, err) {
		return
	}

	fmt.Println("--> Creating 1st Alert on ReplicaSet")
	labelMap := map[string]string{
		"objectType": host.TypeReplicasets,
		"objectName": replicaSet.Name,
	}
	firstAlert, err := mini.CreateAlert(watcher, replicaSet.Namespace, labelMap, host.CheckCommandVolume)
	if !assert.Nil(t, err) {
		return
	}

	// Check Icinga Objects for 1st Alert.
	fmt.Println("----> Checking Icinga Objects for 1st Alert")
	if err := util.CheckIcingaObjectsForAlert(watcher, firstAlert, false, false); !assert.Nil(t, err) {
		return
	}
	fmt.Println("---->> Check Successful")

	fmt.Println("--> Creating 2nd Alert on ReplicaSet")
	secondAlert, err := mini.CreateAlert(watcher, replicaSet.Namespace, labelMap, host.CheckCommandVolume)
	if !assert.Nil(t, err) {
		return
	}

	// Check Icinga Objects for 2nd Alert.
	fmt.Println("----> Checking Icinga Objects for 2nd Alert")
	if err := util.CheckIcingaObjectsForAlert(watcher, secondAlert, false, false); !assert.Nil(t, err) {
		return
	}
	fmt.Println("---->> Check Successful")

	// Increment Replica
	fmt.Println("--> Incrementing Replica")
	replicaSet.Spec.Replicas++
	if replicaSet, err = mini.UpdateReplicaSet(watcher, replicaSet); !assert.Nil(t, err) {
		return
	}

	// Get Last Replica
	fmt.Println("--> Getting Last Replica")
	lastPod, err := mini.GetLastReplica(watcher, replicaSet)
	if !assert.Nil(t, err) {
		return
	}

	// Checking Icinga Objects for This Pod
	fmt.Println("----> Checking Icinga Objects for last Pod")
	if err = util.CheckIcingaObjectsForPod(watcher, lastPod.Name, lastPod.Namespace, 2); !assert.Nil(t, err) {
		return
	}
	fmt.Println("---->> Check Successful")


	// Delete 1st Alert
	fmt.Println("--> Deleting 1st Alert")
	if err := mini.DeleteAlert(watcher, firstAlert); !assert.Nil(t, err) {
		return
	}

	// Check Icinga Objects for 2nd Alert.
	fmt.Println("----> Checking Icinga Objects for 1st Alert")
	if err := util.CheckIcingaObjectsForAlert(watcher, firstAlert, false, true); !assert.Nil(t, err) {
		return
	}
	fmt.Println("---->> Check Successful")

	// Delete 2nd Alert
	fmt.Println("--> Deleting 2nd Alert")
	if err := mini.DeleteAlert(watcher, secondAlert); !assert.Nil(t, err) {
		return
	}

	// Check Icinga Objects for 2nd Alert.
	fmt.Println("----> Checking Icinga Objects for 2nd Alert")
	if err := util.CheckIcingaObjectsForAlert(watcher, secondAlert, true, true); !assert.Nil(t, err) {
		return
	}
	fmt.Println("---->> Check Successful")

	// Delete ReplicaSet
	fmt.Println("--> Deleting ReplicaSet")
	if err := mini.DeleteReplicaSet(watcher, replicaSet); !assert.Nil(t, err) {
		return
	}
}

func TestMultipleAlertsOnMultipleObjects(t *testing.T) {
	// Run KubeD
	// runKubeD(setIcingaClient bool)
	// Pass true to set IcingaClient in watcher
	watcher, err := runKubeD(true)
	if !assert.Nil(t, err) {
		return
	}
	fmt.Println("--> Running kubeD")

	// Create ReplicaSet
	fmt.Println("--> Creating ReplicaSet")
	replicaSet, err := mini.CreateReplicaSet(watcher, "default")
	if !assert.Nil(t, err) {
		return
	}

	// Create Service on ReplicaSet
	fmt.Println("--> Creating Service")
	service, err := mini.CreateService(watcher, replicaSet.Namespace, replicaSet.Spec.Template.Labels)
	if !assert.Nil(t, err) {
		return
	}

	// Create 1st Alert
	fmt.Println("--> Creating 1st Alert on ReplicaSet")
	labelMap := map[string]string{
		"objectType": host.TypeReplicasets,
		"objectName": replicaSet.Name,
	}
	firstAlert, err := mini.CreateAlert(watcher, replicaSet.Namespace, labelMap, host.CheckCommandVolume)
	if !assert.Nil(t, err) {
		return
	}

	// Check Icinga Objects for 1st Alert.
	fmt.Println("----> Checking Icinga Objects for 1st Alert")
	if err := util.CheckIcingaObjectsForAlert(watcher, firstAlert, false, false); !assert.Nil(t, err) {
		return
	}
	fmt.Println("---->> Check Successful")

	// Create 2nd Alert
	fmt.Println("--> Creating 2nd Alert on Service")
	labelMap = map[string]string{
		"objectType": host.TypeServices,
		"objectName": service.Name,
	}
	secondAlert, err := mini.CreateAlert(watcher, service.Namespace, labelMap, host.CheckCommandVolume)
	if !assert.Nil(t, err) {
		return
	}

	// Check Icinga Objects for 2nd Alert.
	fmt.Println("----> Checking Icinga Objects for 2nd Alert")
	if err := util.CheckIcingaObjectsForAlert(watcher, secondAlert, false, false); !assert.Nil(t, err) {
		return
	}
	fmt.Println("---->> Check Successful")

	// Get Pod
	fmt.Println("--> Getting Pod")
	pod, err := mini.GetLastReplica(watcher, replicaSet)
	if !assert.Nil(t, err) {
		return
	}

	// Checking Icinga Objects for Pod
	fmt.Println("----> Checking Icinga Objects for Pod")
	if err = util.CheckIcingaObjectsForPod(watcher, pod.Name, pod.Namespace, 2); !assert.Nil(t, err) {
		return
	}
	fmt.Println("---->> Check Successful")

	// Delete Pod
	fmt.Println("--> Deleting Pod")
	if err := mini.DeletePod(watcher, pod); !assert.Nil(t, err) {
		return
	}

	// Checking Icinga Objects for Pod
	fmt.Println("----> Checking Icinga Objects for Pod")
	if err = util.CheckIcingaObjectsForPod(watcher, pod.Name, pod.Namespace, 0); !assert.Nil(t, err) {
		return
	}
	fmt.Println("---->> Check Successful")

	// Getting ReplicaSetObjectList
	replicaSetObjectList, err := util.GetIcingaHostList(watcher, firstAlert)
	if !assert.Nil(t, err) {
		return
	}

	// Delete ReplicaSet
	fmt.Println("--> Deleting ReplicaSet")
	if err := mini.DeleteReplicaSet(watcher, replicaSet); !assert.Nil(t, err) {
		return
	}

	// Check Icinga Objects for 1st Alert.
	fmt.Println("----> Checking Icinga Objects for 1st Alert")
	if err := util.CheckIcingaObjects(watcher, firstAlert, replicaSetObjectList, true, true); !assert.Nil(t, err) {
		return
	}
	fmt.Println("---->> Check Successful")

	// Delete Service
	fmt.Println("--> Deleting Service")
	if err := mini.DeleteService(watcher, service); !assert.Nil(t, err) {
		return
	}

	// Delete 1st Alert
	fmt.Println("--> Deleting 1st Alert")
	if err := mini.DeleteAlert(watcher, firstAlert); !assert.Nil(t, err) {
		return
	}

	// Delete 2nd Alert
	fmt.Println("--> Deleting 2nd Alert")
	if err := mini.DeleteAlert(watcher, secondAlert); !assert.Nil(t, err) {
		return
	}
}
