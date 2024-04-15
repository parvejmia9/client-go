package main

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func main() {
	home, _ := os.UserHomeDir()
	kubeConfigPath := filepath.Join(home, ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	deploymentsClients := clientset.AppsV1().Deployments(v1.NamespaceDefault)
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"app": "bookapi",
			},
			Name: "bookdeployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": "bookapi"},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": "bookapi"},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "bookapi",
							Image: "parvejmia9/api-server:0.0.4",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 9090,
								},
							},
						},
					},
				},
			},
		},
	}
	fmt.Println("Creating deployment...")
	result, err := deploymentsClients.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())

	serviceClients := clientset.CoreV1().Services(v1.NamespaceDefault)
	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "bookservice",
		},
		Spec: apiv1.ServiceSpec{
			Type: v1.ServiceTypeNodePort,
			Selector: map[string]string{
				"app": "bookapi",
			},
			Ports: []apiv1.ServicePort{
				{
					Protocol:   apiv1.ProtocolTCP,
					Port:       3200,
					NodePort:   30001,
					TargetPort: intstr.FromInt(9090),
				},
			},
		},
	}

	fmt.Println("Creating Service...")
	res, err := serviceClients.Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created Service %q.\n", res.GetObjectMeta().GetName())

}

func int32Ptr(i int32) *int32 { return &i }
