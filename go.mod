module github.com/jmainguy/k8sGetGroups

go 1.15

replace k8s.io/client-go => k8s.io/client-go v0.28.0

require (
	github.com/openshift/api v0.0.0-20230816181854-a7ca92db022a
	github.com/openshift/client-go v0.0.0-20230807132528-be5346fb33cb
	github.com/sirupsen/logrus v1.9.3
	k8s.io/apimachinery v0.28.0
	k8s.io/client-go v0.28.0
)
