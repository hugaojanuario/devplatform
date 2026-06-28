package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthRoutes(rg *gin.RouterGroup) {

	health := rg.Group("/health")

	health.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "OK",
		})
	})
}
