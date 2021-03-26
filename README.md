# k8sGetGroups
[![Go Report Card](https://goreportcard.com/badge/github.com/Jmainguy/k8sGetGroups)](https://goreportcard.com/badge/github.com/Jmainguy/k8sGetGroups)
[![Release](https://img.shields.io/github/release/Jmainguy/k8sGetGroups.svg?style=flat-square)](https://github.com/Jmainguy/k8sGetGroups/releases/latest)
[![Coverage Status](https://coveralls.io/repos/github/Jmainguy/k8sGetGroups/badge.svg?branch=main)](https://coveralls.io/github/Jmainguy/k8sGetGroups?branch=main)

Get kubernetes groups from a list of namespaces

# Usage
```/bin/bash
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

## PreBuilt Binaries
Grab Binaries from [The Releases Page](https://github.com/Jmainguy/k8sGetGroups/releases)

## Build
```/bin/bash
export GO111MODULE=on
go build

```
