package plugin

import (
	api "github.com/appscode/searchlight/apis/monitoring/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetComponentStatusPlugin() *api.SearchlightPlugin {
	return &api.SearchlightPlugin{
		ObjectMeta: metav1.ObjectMeta{
			Name: "component-status",
		},
		TypeMeta: PluginTypeMeta,
		Spec: api.SearchlightPluginSpec{
			Command:    "hyperalert check_component_status",
			AlertKinds: []string{api.ResourceKindClusterAlert},
			Arguments: api.PluginArguments{
				Vars: &api.PluginVars{
					Items: map[string]api.PluginVarItem{
						"selector": {
							Type: api.VarTypeString,
						},
						"componentName": {
							Type: api.VarTypeString,
						},
					},
				},
				Host: map[string]string{
					"v": "vars.verbosity",
				},
			},
			State: []string{stateOK, stateCritical, stateUnknown},
		},
	}
}

func GetJsonPathPlugin() *api.SearchlightPlugin {
	return &api.SearchlightPlugin{
		ObjectMeta: metav1.ObjectMeta{
			Name: "json-path",
		},
		TypeMeta: PluginTypeMeta,
		Spec: api.SearchlightPluginSpec{
			Command:    "hyperalert check_json_path",
			AlertKinds: []string{api.ResourceKindClusterAlert},
			Arguments: api.PluginArguments{
				Vars: &api.PluginVars{
					Items: map[string]api.PluginVarItem{
						"url": {
							Type: api.VarTypeString,
						},
						"secretName": {
							Type: api.VarTypeString,
						},
						"warning": {
							Type: api.VarTypeString,
						},
						"critical": {
							Type: api.VarTypeString,
						},
					},
					Required: []string{"url"},
				},
				Host: map[string]string{
					"host": "name",
					"v":    "vars.verbosity",
				},
			},
			State: []string{stateOK, stateWarning, stateCritical, stateUnknown},
		},
	}
}

func GetNodeExistsPlugin() *api.SearchlightPlugin {
	return &api.SearchlightPlugin{
		ObjectMeta: metav1.ObjectMeta{
			Name: "node-exists",
		},
		TypeMeta: PluginTypeMeta,
		Spec: api.SearchlightPluginSpec{
			Command:    "hyperalert check_node_exists",
			AlertKinds: []string{api.ResourceKindClusterAlert},
			Arguments: api.PluginArguments{
				Vars: &api.PluginVars{
					Items: map[string]api.PluginVarItem{
						"selector": {
							Type: api.VarTypeString,
						},
						"nodeName": {
							Type: api.VarTypeString,
						},
						"count": {
							Type: api.VarTypeInteger,
						},
					},
				},
				Host: map[string]string{
					"v": "vars.verbosity",
				},
			},
			State: []string{stateOK, stateCritical, stateUnknown},
		},
	}
}

func GetPodExistsPlugin() *api.SearchlightPlugin {
	return &api.SearchlightPlugin{
		ObjectMeta: metav1.ObjectMeta{
			Name: "pod-exists",
		},
		TypeMeta: PluginTypeMeta,
		Spec: api.SearchlightPluginSpec{
			Command:    "hyperalert check_pod_exists",
			AlertKinds: []string{api.ResourceKindClusterAlert},
			Arguments: api.PluginArguments{
				Vars: &api.PluginVars{
					Items: map[string]api.PluginVarItem{
						"selector": {
							Type: api.VarTypeString,
						},
						"podName": {
							Type: api.VarTypeString,
						},
						"count": {
							Type: api.VarTypeInteger,
						},
					},
				},
				Host: map[string]string{
					"host": "name",
					"v":    "vars.verbosity",
				},
			},
			State: []string{stateOK, stateCritical, stateUnknown},
		},
	}
}

func GetEventPlugin() *api.SearchlightPlugin {
	return &api.SearchlightPlugin{
		ObjectMeta: metav1.ObjectMeta{
			Name: "event",
		},
		TypeMeta: PluginTypeMeta,
		Spec: api.SearchlightPluginSpec{
			Command:    "hyperalert check_event",
			AlertKinds: []string{api.ResourceKindClusterAlert},
			Arguments: api.PluginArguments{
				Vars: &api.PluginVars{
					Items: map[string]api.PluginVarItem{
						"clockSkew": {
							Type: api.VarTypeDuration,
						},
						"involvedObjectName": {
							Type: api.VarTypeString,
						},
						"involvedObjectNamespace": {
							Type: api.VarTypeString,
						},
						"involvedObjectKind": {
							Type: api.VarTypeString,
						},
						"involvedObjectUID": {
							Type: api.VarTypeString,
						},
					},
				},
				Host: map[string]string{
					"host": "name",
					"v":    "vars.verbosity",
				},
			},
			State: []string{stateOK, stateWarning, stateUnknown},
		},
	}
}

func GetCACertPlugin() *api.SearchlightPlugin {
	return &api.SearchlightPlugin{
		ObjectMeta: metav1.ObjectMeta{
			Name: "ca-cert",
		},
		TypeMeta: PluginTypeMeta,
		Spec: api.SearchlightPluginSpec{
			Command:    "hyperalert check_ca_cert",
			AlertKinds: []string{api.ResourceKindClusterAlert},
			Arguments: api.PluginArguments{
				Vars: &api.PluginVars{
					Items: map[string]api.PluginVarItem{
						"warning": {
							Type: api.VarTypeDuration,
						},
						"critical": {
							Type: api.VarTypeDuration,
						},
					},
				},
				Host: map[string]string{
					"v": "vars.verbosity",
				},
			},
			State: []string{stateOK, stateWarning, stateCritical, stateUnknown},
		},
	}
}

func GetCertPlugin() *api.SearchlightPlugin {
	return &api.SearchlightPlugin{
		ObjectMeta: metav1.ObjectMeta{
			Name: "cert",
		},
		TypeMeta: PluginTypeMeta,
		Spec: api.SearchlightPluginSpec{
			Command:    "hyperalert check_cert",
			AlertKinds: []string{api.ResourceKindClusterAlert},
			Arguments: api.PluginArguments{
				Vars: &api.PluginVars{
					Items: map[string]api.PluginVarItem{
						"selector": {
							Type: api.VarTypeString,
						},
						"secretName": {
							Type: api.VarTypeString,
						},
						"secretKey": {
							Type: api.VarTypeString,
						},
						"warning": {
							Type: api.VarTypeDuration,
						},
						"critical": {
							Type: api.VarTypeDuration,
						},
					},
				},
				Host: map[string]string{
					"host": "name",
					"v":    "vars.verbosity",
				},
			},
			State: []string{stateOK, stateWarning, stateCritical, stateUnknown},
		},
	}
}
