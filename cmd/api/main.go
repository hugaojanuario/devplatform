package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hugaojanuario/devplatform/internal/http"
	"github.com/hugaojanuario/devplatform/internal/logger"
	"github.com/hugaojanuario/devplatform/pkg/config"
	database "github.com/hugaojanuario/devplatform/pkg/database/postgres"
)

func main() {
	logs := logger.Logger()
	cfg := config.LoadEnvFile()

	db, err := database.Conn(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	r := gin.Default()
	http.ApiRoutes(r)
	r.Run(":" + cfg.Api.Port)

	logs.Info("API running")
}
