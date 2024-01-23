package main

import (
	"context"
	"fmt"

	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pointer"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/utils/pointer"
	metav1 "k8s.ip/apimachinery/pkg/apis/meta/v1"
)

const (
	NAMESPACE  = "dev"
	DEPLOYMENT = "dev-deployment"
	SERVICE    = "dev-service"
)

func initConfig() (*rest.Config, error) {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func initClient() (*kubernetes.Clientset, error) {
	config, err := initConfig()
	if err != nil {
		panic(err)
	}
	return kubernetes.NewForConfig(config)
}

func createNamespace(clientSet *kubernetes.Clientset) {
	nsClient := clientSet.CoreV1().Namespaces()
	namespace := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: NAMESPACE,
		},
		Status: corev1.NamespaceStatus{},
	}
	ns, err := nsClient.Create(context.TODO(), namespace, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(ns.GetName())
}

func createDeployment(clientSet *kubernetes.Clientset) {
	deployClient := clientSet.AppsV1().Deployments(NAMESPACE)
	deployment := &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      DEPLOYMENT,
			Namespace: NAMESPACE,
		},
		Spec: v1.DeploymentSpec{
			Replicas: pointer.Int32(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "nginx",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: "nginx",
					Labels: map[string]string{
						"app": "nginx",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx",
							Image: "nginx:1.17.1",
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 80,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							ImagePullPolicy: "IfNotPresent",
						},
					},
				},
			},
		},
		Status: v1.DeploymentStatus{},
	}
	result, err := deployClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(result.GetName())
}

func createService(clientSet *kubernetes.Clientset) {
	serviceClient := clientSet.CoreV1().Services(NAMESPACE)
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      SERVICE,
			Namespace: NAMESPACE,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Port:     80,
					NodePort: 30001,
				},
			},
			Selector: map[string]string{
				"app": "nginx",
			},
			Type: corev1.ServiceTypeNodePort,
		},
		Status: corev1.ServiceStatus{},
	}
	result, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(result.GetName())

}

func deleteAll() {
	clientSet, err := initClient()
	if err != nil {
		return
	}

	err = clientSet.CoreV1().Services(NAMESPACE).Delete(context.TODO(), SERVICE, metav1.DeleteOptions{})
	if err != nil {
		return
	}

	err = clientSet.AppsV1().Deployments(NAMESPACE).Delete(context.TODO(), DEPLOYMENT, metav1.DeleteOptions{})
	if err != nil {
		return
	}

	err = clientSet.CoreV1().Namespaces().Delete(context.TODO(), NAMESPACE, metav1.DeleteOptions{})
	if err != nil {
		return
	}

}

func main() {

}
