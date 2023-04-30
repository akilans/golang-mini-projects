// This application
// Lists Pods, Deployments in the kubernetes cluster
// Creates/Updates a deployment with k8s manifest yaml file

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
)

// Function to list all the pods in the specific namespace
func listPods(clientset *kubernetes.Clientset, namespace string) {

	pods, err := clientset.CoreV1().Pods(namespace).List(context.Background(), v1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	// print all the pod names in the specified namespace
	if len(pods.Items) > 0 {
		for _, pod := range pods.Items {
			fmt.Println("Pod -", pod.Name)
		}
	} else {
		fmt.Printf("No pods found in %s namespace\n", namespace)
	}
}

// Function to list all the Deployments in the specific namespace
func listDeployments(clientset *kubernetes.Clientset, namespace string) {

	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.Background(), v1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	// print all the pod names in default namespace
	if len(deployments.Items) > 0 {
		for _, deployment := range deployments.Items {
			fmt.Println("Deployment -", deployment.Name)
		}
	} else {
		fmt.Printf("No deployments found in %s namespace\n", namespace)
	}
}

func main() {
	fmt.Println("Testing client go...")

	// Namespace to deploy
	namespace := "default"
	kubeConfigFilePath := "/home/akilan/.kube/config"

	kubeconfig := flag.String("kubeconfig", kubeConfigFilePath, "kubeconfig file location")

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// getting pods in default namespace
	listPods(clientset, namespace)
	// get all the deployments in the default namespace
	listDeployments(clientset, namespace)

	// Load the file into a buffer
	data, err := os.ReadFile("dep.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Universal Decoder for schema decoding for all kinds [ Pod, Deployment, Service etc...]
	decoder := scheme.Codecs.UniversalDeserializer()

	// Find GVK and decoding
	obj, groupVersionKind, err := decoder.Decode(data, nil, nil)
	if err != nil {
		fmt.Println(err)
	}

	// Supporting only deployment kind
	if groupVersionKind.Group == "apps" &&
		groupVersionKind.Version == "v1" &&
		groupVersionKind.Kind == "Deployment" {

		deploymentObj := obj.(*appsv1.Deployment)

		// check if deployment with same name exists
		existingDeployments, _ := clientset.AppsV1().Deployments(namespace).List(context.Background(), v1.ListOptions{})
		if err != nil {
			fmt.Println(err)
		}

		isDeploymentExists := false

		for _, dep := range existingDeployments.Items {
			if dep.Name == deploymentObj.ObjectMeta.Name {
				isDeploymentExists = true
			}
		}

		if !isDeploymentExists {
			// create a deployment if empty
			fmt.Printf("%s is not found in %s namespace. So Creating...\n", deploymentObj.ObjectMeta.Name, namespace)
			_, err = clientset.AppsV1().Deployments(namespace).Create(context.Background(), deploymentObj, v1.CreateOptions{})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Created deployment successfully...")
			}

		} else {
			// update the deployment if exists
			fmt.Printf("%s is found in %s namespace. So Updating.... \n", deploymentObj.ObjectMeta.Name, namespace)
			_, err = clientset.AppsV1().Deployments(namespace).Update(context.Background(), deploymentObj, v1.UpdateOptions{})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Updated deployment successfully...")
			}
		}

	}

}
