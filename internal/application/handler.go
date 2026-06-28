package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	s *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{s: s}
}

func (h *Handler) Create(ctx *gin.Context) {
	var req CreateApplication
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	app, err := h.s.Create(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, app)
}

func (h *Handler) FindAll(ctx *gin.Context) {
	apps, err := h.s.FindAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, apps)
}

func (h *Handler) FindById(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id invalid"})
		return
	}

	app, err := h.s.FindByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "id invalid"})
		return
	}

	ctx.JSON(http.StatusOK, app)
}

func (h *Handler) Update(ctx *gin.Context) {
	var req UpdateApplication
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id invalid"})
		return
	}

	app, err := h.s.Update(ctx, id, req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "id invalid"})
		return
	}

	ctx.JSON(http.StatusOK, app)
}

func (h *Handler) Delete(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id invalid"})
		return
	}

	if err := h.s.Delete(ctx, id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id invalid"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
