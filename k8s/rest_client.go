package main

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"

	client, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	result := &v1.PodList{}
	namespace := "kube-system"
	err = client.Get().Namespace(namespace).Resource("pods").VersionedParams(&metav1.ListOptions{Limit: 100}, scheme.ParameterCodec).Do(context.TODO()).Into(result)
	if err != nil {
		panic(err)
	}

	for _, pod := range result.Items {
		fmt.Println(pod.Name)
	}
}
