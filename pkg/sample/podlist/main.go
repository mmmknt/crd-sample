package main

import (
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String(
		"kubeconfig",
		"/Users/asato/.kube/config",
		"kubeconfig config file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	for i, pod := range pods.Items {
		fmt.Printf("[Pod Name %d]%s\n", i, pod.GetName())
	}


}
