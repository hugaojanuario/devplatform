package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {
	health := r.Group("health")

	health.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "OK",
		})
	})
}
