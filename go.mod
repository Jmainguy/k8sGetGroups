module github.com/jmainguy/k8sGetGroups

go 1.15

replace k8s.io/client-go => k8s.io/client-go v0.27.4

require (
	github.com/openshift/api v0.0.0-20230815081605-d7f5a4c98505
	github.com/openshift/client-go v0.0.0-20230807132528-be5346fb33cb
	github.com/sirupsen/logrus v1.9.3
	k8s.io/apimachinery v0.27.4
	k8s.io/client-go v0.27.4
)
