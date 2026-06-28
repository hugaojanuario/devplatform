package application

import "github.com/gin-gonic/gin"

func ApplicationRoutes(rg *gin.RouterGroup, h *Handler) {

	applications := rg.Group("/applications")

	applications.GET("/", h.FindAll)
	applications.GET("/:id", h.FindById)
	applications.POST("/", h.Create)
	applications.PUT("/:id", h.Update)
	applications.DELETE("/:id", h.Delete)
}
