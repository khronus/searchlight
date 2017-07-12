package check_volume

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/appscode/go/flags"
	"github.com/appscode/go/net/httpclient"
	"github.com/appscode/searchlight/pkg/icinga"
	"github.com/appscode/searchlight/pkg/util"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiv1 "k8s.io/client-go/pkg/api/v1"
)

const (
	awsElasticBlockStorePluginName = "kubernetes.io~aws-ebs"
	azureDataDiskPluginName        = "kubernetes.io~azure-disk"
	azureFilePluginName            = "kubernetes.io~azure-file"
	cephfsPluginName               = "kubernetes.io~cephfs"
	cinderVolumePluginName         = "kubernetes.io~cinder"
	configMapPluginName            = "kubernetes.io~configmap"
	downwardAPIPluginName          = "kubernetes.io~downward-api"
	emptyDirPluginName             = "kubernetes.io~empty-dir"
	fcPluginName                   = "kubernetes.io~fc"
	flockerPluginName              = "kubernetes.io~flocker"
	gcePersistentDiskPluginName    = "kubernetes.io~gce-pd"
	gitRepoPluginName              = "kubernetes.io~git-repo"
	glusterfsPluginName            = "kubernetes.io~glusterfs"
	hostPathPluginName             = "kubernetes.io~host-path"
	iscsiPluginName                = "kubernetes.io~iscsi"
	nfsPluginName                  = "kubernetes.io~nfs"
	quobytePluginName              = "kubernetes.io~quobyte"
	rbdPluginName                  = "kubernetes.io~rbd"
	secretPluginName               = "kubernetes.io~secret"
	vsphereVolumePluginName        = "kubernetes.io~vsphere-volume"
)

func getVolumePluginName(volumeSource *apiv1.VolumeSource) string {
	if volumeSource.AWSElasticBlockStore != nil {
		return awsElasticBlockStorePluginName
	} else if volumeSource.AzureDisk != nil {
		return azureDataDiskPluginName
	} else if volumeSource.AzureFile != nil {
		return azureFilePluginName
	} else if volumeSource.CephFS != nil {
		return cephfsPluginName
	} else if volumeSource.Cinder != nil {
		return cinderVolumePluginName
	} else if volumeSource.ConfigMap != nil {
		return configMapPluginName
	} else if volumeSource.DownwardAPI != nil {
		return downwardAPIPluginName
	} else if volumeSource.EmptyDir != nil {
		return emptyDirPluginName
	} else if volumeSource.FC != nil {
		return fcPluginName
	} else if volumeSource.Flocker != nil {
		return flockerPluginName
	} else if volumeSource.GCEPersistentDisk != nil {
		return gcePersistentDiskPluginName
	} else if volumeSource.GitRepo != nil {
		return gitRepoPluginName
	} else if volumeSource.Glusterfs != nil {
		return glusterfsPluginName
	} else if volumeSource.HostPath != nil {
		return hostPathPluginName
	} else if volumeSource.ISCSI != nil {
		return iscsiPluginName
	} else if volumeSource.NFS != nil {
		return nfsPluginName
	} else if volumeSource.Quobyte != nil {
		return quobytePluginName
	} else if volumeSource.RBD != nil {
		return rbdPluginName
	} else if volumeSource.Secret != nil {
		return secretPluginName
	} else if volumeSource.VsphereVolume != nil {
		return vsphereVolumePluginName
	}
	return ""
}

func getPersistentVolumePluginName(volumeSource *apiv1.PersistentVolumeSource) string {
	if volumeSource.AWSElasticBlockStore != nil {
		return awsElasticBlockStorePluginName
	} else if volumeSource.AzureDisk != nil {
		return azureDataDiskPluginName
	} else if volumeSource.AzureFile != nil {
		return azureFilePluginName
	} else if volumeSource.CephFS != nil {
		return cephfsPluginName
	} else if volumeSource.Cinder != nil {
		return cinderVolumePluginName
	} else if volumeSource.FC != nil {
		return fcPluginName
	} else if volumeSource.Flocker != nil {
		return flockerPluginName
	} else if volumeSource.GCEPersistentDisk != nil {
		return gcePersistentDiskPluginName
	} else if volumeSource.Glusterfs != nil {
		return glusterfsPluginName
	} else if volumeSource.HostPath != nil {
		return hostPathPluginName
	} else if volumeSource.ISCSI != nil {
		return iscsiPluginName
	} else if volumeSource.NFS != nil {
		return nfsPluginName
	} else if volumeSource.Quobyte != nil {
		return quobytePluginName
	} else if volumeSource.RBD != nil {
		return rbdPluginName
	} else if volumeSource.VsphereVolume != nil {
		return vsphereVolumePluginName
	}
	return ""
}

const (
	hostFactPort = 56977
)

type Request struct {
	Host       string
	NodeStat   bool
	SecretName string
	VolumeName string
	Warning    float64
	Critical   float64
}

type usageStat struct {
	Path              string  `json:"path"`
	Fstype            string  `json:"fstype"`
	Total             uint64  `json:"total"`
	Free              uint64  `json:"free"`
	Used              uint64  `json:"used"`
	UsedPercent       float64 `json:"usedPercent"`
	InodesTotal       uint64  `json:"inodesTotal"`
	InodesUsed        uint64  `json:"inodesUsed"`
	InodesFree        uint64  `json:"inodesFree"`
	InodesUsedPercent float64 `json:"inodesUsedPercent"`
}

type authInfo struct {
	ca        []byte
	key       []byte
	crt       []byte
	authToken string
	username  string
	password  string
}

const (
	ca        = "ca.crt"
	key       = "hostfacts.key"
	crt       = "hostfacts.crt"
	authToken = "auth_token"
	username  = "username"
	password  = "password"
)

func getHostfactsSecretData(kubeClient *util.KubeClient, secretName, secretNamespace string) *authInfo {
	if secretName == "" {
		return nil
	}

	secret, err := kubeClient.Client.CoreV1().Secrets(secretNamespace).Get(secretName, metav1.GetOptions{})
	if err != nil {
		return nil
	}

	authData := &authInfo{
		ca:        secret.Data[ca],
		key:       secret.Data[key],
		crt:       secret.Data[crt],
		authToken: string(secret.Data[authToken]),
		username:  string(secret.Data[username]),
		password:  string(secret.Data[password]),
	}

	return authData
}

func getUsage(authInfo *authInfo, hostIP, path string) (*usageStat, error) {
	scheme := "http"
	httpClient := httpclient.Default()
	if authInfo != nil && authInfo.ca != nil {
		scheme = "https"
		httpClient.WithBasicAuth(authInfo.username, authInfo.password).
			WithBearerToken(authInfo.authToken).
			WithTLSConfig(authInfo.ca, authInfo.crt, authInfo.key)
	}

	urlStr := fmt.Sprintf("%v://%v:%v/du?p=%v", scheme, hostIP, hostFactPort, path)
	usages := make([]*usageStat, 1)
	_, err := httpClient.Call(http.MethodGet, urlStr, nil, &usages, true)
	if err != nil {
		return nil, err
	}

	return usages[0], nil
}

func checkResult(field string, warning, critical, result float64) (icinga.State, interface{}) {
	if result >= critical {
		return icinga.CRITICAL, fmt.Sprintf("%v used more than %v%%", field, critical)
	}
	if result >= warning {
		return icinga.WARNING, fmt.Sprintf("%v used more than %v%%", field, warning)
	}
	return icinga.OK, "(Disk & Inodes)"
}

func checkDiskStat(kubeClient *util.KubeClient, req *Request, namespace, nodeIP, path string) (icinga.State, interface{}) {

	authInfo := getHostfactsSecretData(kubeClient, req.SecretName, namespace)
	usage, err := getUsage(authInfo, nodeIP, path)
	if err != nil {
		return icinga.UNKNOWN, err
	}

	warning := req.Warning
	critical := req.Critical
	state, message := checkResult("Disk", warning, critical, usage.UsedPercent)
	if state != icinga.OK {
		return state, message
	}
	state, message = checkResult("Inodes", warning, critical, usage.InodesUsedPercent)
	return state, message
}

func checkNodeDiskStat(req *Request) (icinga.State, interface{}) {
	host, err := icinga.ParseHost(req.Host)
	if err != nil {
		return icinga.UNKNOWN, "Invalid icinga host.name"
	}
	if host.Type != icinga.TypeNode {
		return icinga.UNKNOWN, "Invalid icinga host type"
	}

	kubeClient, err := util.NewClient()
	if err != nil {
		return icinga.UNKNOWN, err
	}

	node, err := kubeClient.Client.CoreV1().Nodes().Get(host.ObjectName, metav1.GetOptions{})
	if err != nil {
		return icinga.UNKNOWN, err
	}

	if node == nil {
		return icinga.UNKNOWN, "Node not found"
	}

	hostIP := ""
	for _, address := range node.Status.Addresses {
		if address.Type == apiv1.NodeInternalIP {
			hostIP = address.Address
		}
	}

	if hostIP == "" {
		return icinga.UNKNOWN, "Node InternalIP not found"
	}
	return checkDiskStat(kubeClient, req, host.AlertNamespace, hostIP, "/")
}

func checkPodVolumeStat(req *Request) (icinga.State, interface{}) {
	host, err := icinga.ParseHost(req.Host)
	if err != nil {
		return icinga.UNKNOWN, "Invalid icinga host.name"
	}
	if host.Type != icinga.TypePod {
		return icinga.UNKNOWN, "Invalid icinga host type"
	}


	kubeClient, err := util.NewClient()
	if err != nil {
		return icinga.UNKNOWN, err
	}

	pod, err := kubeClient.Client.CoreV1().Pods(host.AlertNamespace).Get(host.ObjectName, metav1.GetOptions{})

	if err != nil {
		return icinga.UNKNOWN, err
	}

	var volumeSourcePluginName = ""
	var volumeSourceName = ""
	for _, volume := range pod.Spec.Volumes {
		if volume.Name == req.VolumeName {
			if volume.PersistentVolumeClaim != nil {

				claim, err := kubeClient.Client.CoreV1().PersistentVolumeClaims(host.AlertNamespace).Get(volume.PersistentVolumeClaim.ClaimName, metav1.GetOptions{})
				if err != nil {
					return icinga.UNKNOWN, err

				}
				volume, err := kubeClient.Client.CoreV1().PersistentVolumes().Get(claim.Spec.VolumeName, metav1.GetOptions{})
				if err != nil {
					return icinga.UNKNOWN, err
				}
				volumeSourcePluginName = getPersistentVolumePluginName(&volume.Spec.PersistentVolumeSource)
				volumeSourceName = volume.Name

			} else {
				volumeSourcePluginName = getVolumePluginName(&volume.VolumeSource)
				volumeSourceName = volume.Name
			}
			break
		}
	}

	if volumeSourcePluginName == "" {
		return icinga.UNKNOWN, errors.New("Invalid volume source")
	}

	path := fmt.Sprintf("/var/lib/kubelet/pods/%v/volumes/%v/%v", pod.UID, volumeSourcePluginName, volumeSourceName)
	return checkDiskStat(kubeClient, req, host.AlertNamespace, pod.Status.HostIP, path)
}

func NewCmd() *cobra.Command {
	var req Request
	var icingaHost string

	c := &cobra.Command{
		Use:     "check_volume",
		Short:   "Check kubernetes volume",
		Example: "",

		Run: func(cmd *cobra.Command, args []string) {
			flags.EnsureRequiredFlags(cmd, "host")


			if req.NodeStat {
				icinga.Output(checkNodeDiskStat(&req))
			} else {
				flags.EnsureRequiredFlags(cmd, "volume_name")
				icinga.Output(checkPodVolumeStat(&req))
			}
		},
	}

	c.Flags().StringVarP(&icingaHost, "host", "H", "", "Icinga host name")
	c.Flags().BoolVar(&req.NodeStat, "node_stat", false, "Checking Node disk size")
	c.Flags().StringVarP(&req.SecretName, "secret", "s", "", `Kubernetes secret name`)
	c.Flags().StringVarP(&req.VolumeName, "volume_name", "N", "", "Volume name")
	c.Flags().Float64VarP(&req.Warning, "warning", "w", 75.0, "Warning level value (usage percentage)")
	c.Flags().Float64VarP(&req.Critical, "critical", "c", 90.0, "Critical level value (usage percentage)")
	return c
}
