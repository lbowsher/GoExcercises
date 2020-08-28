package main

import (
	"context"
	"fmt"
	"time"
	"strings"
	// "flag"
	// "path/filepath"

	"github.com/Pallinder/go-randomdata"

	core "k8s.io/api/core/v1"
	//"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	// "k8s.io/client-go/tools/clientcmd"
	// "k8s.io/client-go/util/homedir"
)

func main() {
	// get the kube config needed for creating the clientset
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// loops every 15 seconds
	for {
		// get pods in all the namespaces by omitting namespace
		// Or specify namespace to get pods in particular namespace
		pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		// prints the number of pods
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// prints each pod's name and date of birth
		for _, pod := range pods.Items {
			fmt.Printf("%s %s\n", pod.GetName(), pod.GetCreationTimestamp())
		}

		// build the pod defination we want to deploy
		pod := makePodObject()

		// now create the pod in kubernetes cluster using the clientset
		pod, err = clientset.CoreV1().Pods(pod.Namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
		if err != nil {
			panic(err)
		}
		fmt.Println("Pod created successfully...")

		time.Sleep(15 * time.Second)
	}
}

// returns a random lowercase first name
func getPodName() string{
	return strings.ToLower(randomdata.FirstName(randomdata.RandomGender))
}

// returns the new pod's definition
func makePodObject() *core.Pod {
	podName := getPodName()
	return &core.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: "default",
			Labels: map[string]string{
				"app": "demo",
			},
		},
		Spec: core.PodSpec{
			Containers: []core.Container{
				{
					Name:            "busybox",
					Image:           "busybox",
					ImagePullPolicy: core.PullIfNotPresent,
					Command: []string{
						"sleep",
						"3600",
					},
				},
			},
		},
	}
}
