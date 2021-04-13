package main

import (
	//routev1 "github.com/openshift/api/route/v1"
	"bufio"
	"context"
	"flag"
	routev1 "github.com/openshift/client-go/route/clientset/versioned/typed/route/v1"
	"os"
	"path/filepath"

	"fmt"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	// Support gcp and other authentication schemes
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func getNamespaceListFromFile(namespaceList string) (namespaces []string) {
	file, err := os.Open(namespaceList)

	if err != nil {
		log.Fatalf("failed to open")

	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		namespaces = append(namespaces, scanner.Text())
	}
	file.Close()
	return namespaces

}

func getNamespacesThatHaveRoutes(routeV1Client *routev1.RouteV1Client, clientset *kubernetes.Clientset) (namespacesRoutes []string) {
	// Get all namespaces
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	check(err)

	// Loop through all the namespaces, and see if they have a ingress
	for _, namespace := range namespaces.Items {
		routes, _ := routeV1Client.Routes(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		check(err)
		if len(routes.Items) > 0 {
			namespacesRoutes = append(namespacesRoutes, namespace.Name)
		}
	}
	return
}

func getAdminRoleBindings(clientset *kubernetes.Clientset, namespace string) (adminRoleBindings []string) {
	rolebindings, err := clientset.RbacV1().RoleBindings(namespace).List(context.TODO(), metav1.ListOptions{})
	check(err)
	for _, rolebinding := range rolebindings.Items {
		if rolebinding.RoleRef.Name == "admin" {
			adminRoleBindings = append(adminRoleBindings, rolebinding.Name)
		}
	}
	return
}

func roleBindingsFromNamespaces(clientset *kubernetes.Clientset, namespaces []string) (teams map[string]bool) {
	teams = make(map[string]bool)
	for _, namespace := range namespaces {

		adminRoleBindings := getAdminRoleBindings(clientset, namespace)
		for _, adminRoleBinding := range adminRoleBindings {
			rolebinding, err := clientset.RbacV1().RoleBindings(namespace).Get(context.TODO(), adminRoleBinding, metav1.GetOptions{})
			check(err)
			for _, subject := range rolebinding.Subjects {
				if subject.Kind == "Group" {
					_, exists := teams[subject.Name]
					if !exists {
						teams[subject.Name] = true
					}
				}
			}
		}

	}
	return
}

func routesAndRolebindings(clientset *kubernetes.Clientset, routeV1Client *routev1.RouteV1Client, namespaces []string) {
	for _, namespace := range namespaces {
		var namespacesRoutes []string
		routes, err := routeV1Client.Routes(namespace).List(context.TODO(), metav1.ListOptions{})
		check(err)
		if len(routes.Items) > 0 {
			for _, route := range routes.Items {
				namespacesRoutes = append(namespacesRoutes, route.Spec.Host)
			}
		}
		teams := make(map[string]bool)

		adminRoleBindings := getAdminRoleBindings(clientset, namespace)
		for _, adminRoleBinding := range adminRoleBindings {
			rolebinding, err := clientset.RbacV1().RoleBindings(namespace).Get(context.TODO(), adminRoleBinding, metav1.GetOptions{})
			check(err)
			for _, subject := range rolebinding.Subjects {
				if subject.Kind == "Group" {
					_, exists := teams[subject.Name]
					if !exists {
						teams[subject.Name] = true
					}
				}
			}
		}
		// Ignore namespaces that dont have a group assigned admin
		if len(teams) > 0 {
			//for k, _ := range teams {
			//	roleBindings = fmt.Sprintf("%s '%s'", roleBindings, k)
			//}
			var teamNames []string
			for teamName := range teams {
				teamNames = append(teamNames, teamName)
			}
			msg := fmt.Sprintf("Namespace: %s, Routes %s, Rolebindings %s", namespace, namespacesRoutes, teamNames)
			fmt.Println(msg)
		}
	}
}
func main() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	nameSpace := flag.String("namespace", "", "(optional) Namespace to grab capacity usage from")
	namespaceList := flag.String("namespaceList", "", "(optional) Filepath containing a list of namespaces, one per line")
	showRoutes := flag.Bool("show-routes", false, "Print namespace and routes along with rolebindings")
	checkMode := flag.Bool("check", false, "(optional) Check kubernetes connection")
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		// no config, maybe we are inside a kubernetes cluster.
		config, err = rest.InClusterConfig()
		check(err)
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	check(err)
	// create the route client
	routeV1Client, err := routev1.NewForConfig(config)
	check(err)

	if *checkMode {
		_, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
		check(err)
		fmt.Println("Connection Successful")
		return
	}

	// Get list of namespaces
	var namespaces []string
	if *namespaceList != "" {
		namespaces = getNamespaceListFromFile(*namespaceList)
	} else if *nameSpace != "" {
		namespaces = append(namespaces, *nameSpace)
	} else {
		namespaces = getNamespacesThatHaveRoutes(routeV1Client, clientset)
	}

	if *showRoutes {
		routesAndRolebindings(clientset, routeV1Client, namespaces)
	} else {
		teams := roleBindingsFromNamespaces(clientset, namespaces)
		for teamname := range teams {
			fmt.Println(teamname)
		}
	}
}
