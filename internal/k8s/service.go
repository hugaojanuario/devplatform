package k8s

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type HealthStatus struct {
	Healthy bool
	Message string
}

func CheckHealth() HealthStatus {
	_, err := Client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return HealthStatus{
			Healthy: false,
			Message: err.Error(),
		}
	}

	return HealthStatus{
		Healthy: true,
		Message: "k8s client conectado",
	}
}
