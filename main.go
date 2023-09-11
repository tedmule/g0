package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func main() {

	config := &rest.Config{
		Host:        "https://10.0.0.111:37119",
		BearerToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6IjBKb1Q4Vm40R0dvVjZkZEF1V2gtajIyZnl1RjRYWFRVZ2txSTZOMUVQOGcifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJnYW1lIiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6ImdhbWUtcm9ib3Qtc2VjcmV0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImdhbWUtcm9ib3QiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiJkZjMzMjA2Yi1iMWE2LTQ2M2UtODY4MC1jYzRlZmZiZjRhMGYiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6Z2FtZTpnYW1lLXJvYm90In0.XTv3d6Ks41T6fbKzUHFvzR_ojM5iZ4wRQPQk0Ay-t-ObZrvARZngGCJ2mANEyV9Rc1wKyIU76EQ4IDjy20wwdMmHbiMLE52R_yfc-B7INYVEQDpCKaMwjWIL80f9goewb0YmDLQY5mgNJhJNz8QDQI4vBd2duUscfYNuKlvnD6ceXoURVF_LAdPoFTUN9lmv1ANOrJf_9Bb12OZLAk-Ge-QnqrlJwJCNg_x6QpWqR1mucuGIYbbe0F9TEQY_LZrDyfM0AOOLGPDNbGzQUiLEOtCHWnx1RbH7-wYyi0phwMsj-tzKyLzkZqOKsyDZ-g64PQyNmkgzmYEIJvRGNPBccw",
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	pods, err := clientset.CoreV1().Pods("game").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	for _, v := range pods.Items {
		fmt.Println(v.Name)
	}

	ns, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, v := range ns.Items {
		fmt.Println(v.Name)
	}

}
