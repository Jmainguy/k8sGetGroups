module github.com/jmainguy/k8sGetGroups

go 1.15

replace k8s.io/client-go => k8s.io/client-go v0.27.4

require (
	github.com/openshift/api v0.0.0-20230724190601-61a5301895a5
	github.com/openshift/client-go v0.0.0-20230718165156-6014fb98e86a
	github.com/sirupsen/logrus v1.9.3
	k8s.io/apimachinery v0.27.4
	k8s.io/client-go v0.27.4
)
