package k8s

import (
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var Client *kubernetes.Clientset

func NewClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		home := homedir.HomeDir()
		kubeconfig := filepath.Join(home, ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}

	Client, err = kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return Client, nil
}
