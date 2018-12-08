package k8s

import (
	"testing"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/gruntwork-io/terratest/modules/logger"
)

// GetKubernetesClientFromFileE returns a Kubernetes API client given the kubernetes config file path.
func GetKubernetesClientFromFileE(kubeConfigPath string) (*kubernetes.Clientset, error) {
	// Load API config (instead of more low level ClientConfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

// GetKubernetesClientE returns a Kubernetes API client that can be used to make requests.
func GetKubernetesClientE(t *testing.T) (*kubernetes.Clientset, error) {
	kubeConfigPath, err := GetKubeConfigPathE(t)
	if err != nil {
		return nil, err
	}

	logger.Logf(t, "Configuring kubectl using config file %s", kubeConfigPath)
	return GetKubernetesClientFromFileE(kubeConfigPath)
}
