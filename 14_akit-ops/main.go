package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Common function to check any error
func checkIfError(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
}

func main() {

	fmt.Println("My own git ops - akit-ops")

	// Default values
	url := os.Getenv("REPO_URL")
	pullInterval, _ := strconv.Atoi(os.Getenv("REPO_PULL_INTERVAL"))
	repoPath := "./deployrepo"
	manifestFile := "manifest.yaml"
	namespace := "default"

	// check this app is running inside k8s cluster
	_, isRunningInsideK8sCluster := os.LookupEnv("KUBERNETES_SERVICE_HOST")

	// clientset variable
	var clientset *kubernetes.Clientset

	// create clientset based on where this app is running
	if isRunningInsideK8sCluster {
		// creates the in-cluster config
		config, err := rest.InClusterConfig()
		checkIfError(err)
		// creates the clientset
		clientset, err = kubernetes.NewForConfig(config)
		checkIfError(err)
	} else {
		currentUserHomeDir, _ := os.UserHomeDir()
		kubeConfigFilePath := filepath.Join(currentUserHomeDir, ".kube/config")
		kubeconfig := flag.String("kubeconfig", kubeConfigFilePath, "kubeconfig file location")

		// use the current context in kubeconfig
		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		checkIfError(err)

		// create the clientset
		clientset, err = kubernetes.NewForConfig(config)
		checkIfError(err)
	}

	// call clone repo function
	r := cloneRepo(url, repoPath)

	// Loop to watch for any changes
	for {
		// check for any changes
		isChanged := detectChanges(r)

		if isChanged {

			// Load the file into a buffer
			fname := filepath.Join(repoPath, manifestFile)
			data, err := os.ReadFile(fname)
			checkIfError(err)

			// Universal Decoder for schema decoding for all kinds [ Pod, Deployment, Service etc...]
			decoder := scheme.Codecs.UniversalDeserializer()

			// loop through each kubernetes resources separated by ---
			for _, resourceYAML := range strings.Split(string(data), "---") {

				// skip empty documents
				if len(resourceYAML) == 0 {
					continue
				}

				// Find GVK and decoding.
				obj, groupVersionKind, err := decoder.Decode([]byte(resourceYAML), nil, nil)
				checkIfError(err)

				// check for k8s resource type
				switch groupVersionKind.Kind {
				// deployment
				case "Deployment":
					fmt.Println("Deployment")
					deploymentObj := obj.(*appsv1.Deployment)
					if deploymentObj.ObjectMeta.Namespace != "" {
						namespace = deploymentObj.ObjectMeta.Namespace
					}
					// call the function create/update deployment
					createUpdateDeployment(clientset, deploymentObj, namespace)
				// service
				case "Service":
					fmt.Println("Service")
					serviceObj := obj.(*corev1.Service)
					if serviceObj.ObjectMeta.Namespace != "" {
						namespace = serviceObj.ObjectMeta.Namespace
					}

					// call the function create/update service
					createUpdateService(clientset, serviceObj, namespace)

				default:
					fmt.Printf("Currently %s - resource type is not supported\n", groupVersionKind.Kind)
				}

			}

		}

		fmt.Println("Wait for ", pullInterval, " Seconds")
		time.Sleep(time.Duration(pullInterval) * time.Second)

	}

}
