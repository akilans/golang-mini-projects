// It has function to create/update deployment and Service
package main

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Create / Update service function
func createUpdateService(clientset *kubernetes.Clientset, serviceObj *corev1.Service, namespace string) {

	// check if service with same name exists
	existingservices, err := clientset.CoreV1().Services(namespace).List(context.Background(), v1.ListOptions{})
	checkIfError(err)
	isServiceExists := false

	for _, dep := range existingservices.Items {
		if dep.Name == serviceObj.ObjectMeta.Name {
			isServiceExists = true
		}
	}

	if !isServiceExists {
		// create a service if empty
		fmt.Printf("%s is not found in %s namespace. So Creating...\n", serviceObj.ObjectMeta.Name, namespace)
		_, err := clientset.CoreV1().Services(namespace).Create(context.Background(), serviceObj, v1.CreateOptions{})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Created service successfully...")
		}

	} else {
		// update the service if exists
		fmt.Printf("%s is found in %s namespace. So Updating.... \n", serviceObj.ObjectMeta.Name, namespace)
		_, err := clientset.CoreV1().Services(namespace).Update(context.Background(), serviceObj, v1.UpdateOptions{})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Updated service successfully...")
		}
	}

}

// Create / Update deployment function
func createUpdateDeployment(clientset *kubernetes.Clientset, deploymentObj *appsv1.Deployment, namespace string) {

	// check if deployment with same name exists
	existingDeployments, err := clientset.AppsV1().Deployments(namespace).List(context.Background(), v1.ListOptions{})
	checkIfError(err)
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
