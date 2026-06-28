package aplication

import (
	"time"

	"github.com/google/uuid"
)

type Application struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Namespace string    `json:"namespace"`
	Image     string    `json:"image"`
	Replicas  int       `json:"replicas"`
	Port      int       `json:"port"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateApplication struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Image     string `json:"image"`
	Replicas  int    `json:"replicas"`
	Port      int    `json:"port"`
}

type UpdateApplication struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Image     string `json:"image"`
	Replicas  int    `json:"replicas"`
	Port      int    `json:"port"`
}

type ApplicationResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Namespace string    `json:"namespace"`
	Image     string    `json:"image"`
	Replicas  int       `json:"replicas"`
	Port      int       `json:"port"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
