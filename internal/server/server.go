package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/devplatform/internal/application"
	"github.com/hugaojanuario/devplatform/internal/http"
	"github.com/hugaojanuario/devplatform/internal/k8s"
)

func Server(appHandler *application.Handler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")

	//HEALTH ROUTE
	http.HealthRoutes(api)
	k8s.KubernetesRoutes(api)

	//APPLICATION ROUTES
	application.ApplicationRoutes(api, appHandler)

	return r

}
