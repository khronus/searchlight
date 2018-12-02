---
title: Check Component Status
menu:
  product_searchlight_7.0.0:
    identifier: hyperalert-check-component-status
    name: Check Component Status
    parent: hyperalert-cli
product_name: searchlight
section_menu_id: reference
menu_name: product_searchlight_7.0.0
---
## hyperalert check_component_status

Check Kubernetes Component Status

### Synopsis

Check Kubernetes Component Status

```
hyperalert check_component_status [flags]
```

### Options

```
  -n, --componentName string   Name of component which should be ready
  -h, --help                   help for check_component_status
  -l, --selector string        Selector (label query) to filter on, supports '=', '==', and '!='.
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --bypass-validating-webhook-xray   if true, bypasses validating webhook xray checks
      --context string                   Use the context in kubeconfig
      --icinga.checkInterval int         Icinga check_interval in second. [Format: 30, 300] (default 30)
      --kubeconfig string                Path to kubeconfig file with authorization information (the master location is set by the master flag).
      --log-flush-frequency duration     Maximum number of seconds between log flushes (default 5s)
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files (default true)
      --stderrthreshold severity         logs at or above this threshold go to stderr
      --use-kubeapiserver-fqdn-for-aks   if true, uses kube-apiserver FQDN for AKS cluster to workaround https://github.com/Azure/AKS/issues/522 (default true)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO

* [hyperalert](/docs/reference/hyperalert/hyperalert.md)	 - AppsCode Icinga2 plugin


