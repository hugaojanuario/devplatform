package k8s

import (
	"github.com/gin-gonic/gin"
)

func KubernetesRoutes(rg *gin.RouterGroup) {

	applications := rg.Group("/kubernetes/health")

	applications.GET("", HealthKubernetes)
}
