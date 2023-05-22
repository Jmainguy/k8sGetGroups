module github.com/jmainguy/k8sGetGroups

go 1.15

replace k8s.io/client-go => k8s.io/client-go v0.27.2

require (
	github.com/openshift/api v0.0.0-20210325163602-e37aaed4c278
	github.com/openshift/client-go v0.0.0-20210112165513-ebc401615f47
	github.com/sirupsen/logrus v1.9.2
	k8s.io/apimachinery v0.27.2
	k8s.io/client-go v0.27.2
)
