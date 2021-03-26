# k8sGetGroups
Get kubernetes groups from a list of namespaces

# Usage
```
Usage of k8sGetGroups:
  -check
    	(optional) Check kubernetes connection
  -kubeconfig string
    	(optional) absolute path to the kubeconfig file (default "/home/jmainguy/.kube/config")
  -namespace string
    	(optional) Namespace to grab capacity usage from
  -namespaceList string
    	(optional) Filepath containing a list of namespaces, one per line
```

By default, it will get group names, from all namespaces, that have a openshiftRoute.

If you would prefer a specific namespace, use --namespace

If you would prefer to pass it a list of namespaces, one per line, use --namespaceList

The list returned will be unique group names.
