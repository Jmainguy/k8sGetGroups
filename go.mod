module github.com/jmainguy/k8sGetGroups

go 1.15

replace k8s.io/client-go => k8s.io/client-go v0.27.3

require (
	github.com/openshift/api v0.0.0-20230628121033-34449b12e886
	github.com/openshift/client-go v0.0.0-20210112165513-ebc401615f47
	github.com/sirupsen/logrus v1.9.3
	k8s.io/apimachinery v0.27.3
	k8s.io/client-go v0.27.3
)
