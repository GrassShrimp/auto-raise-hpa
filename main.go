package main

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	deployments, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	for _, deployments := range deployments.Items {
		hpas, err := clientset.AutoscalingV1().HorizontalPodAutoscalers(deployments.Namespace).List(context.TODO(), metav1.ListOptions{})

		if err != nil {
			panic(err.Error())
		}

		for _, hpa := range hpas.Items {
			if deployments.Status.Replicas != hpa.Spec.MaxReplicas {
				continue
			}
			hpa.Spec.MaxReplicas += 10
			go clientset.AutoscalingV1().HorizontalPodAutoscalers(deployments.Namespace).Update(context.TODO(), &hpa, metav1.UpdateOptions{})
		}
	}
}
