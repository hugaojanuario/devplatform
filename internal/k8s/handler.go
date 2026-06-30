package k8s

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func HealthKubernetes(ctx *gin.Context) {

	status := CheckHealth()

	if !status.Healthy {
		ctx.JSON(http.StatusServiceUnavailable, HealthResponse{
			Status:  "unhealthy",
			Message: status.Message,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "OK",
	})
}
