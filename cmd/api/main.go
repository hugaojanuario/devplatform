package main

import (
	"log"

	"github.com/hugaojanuario/devplatform/internal/application"
	"github.com/hugaojanuario/devplatform/internal/logger"
	"github.com/hugaojanuario/devplatform/internal/server"
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

	applicationRepository := application.NewRepository(db)
	applicationService := application.NewService(applicationRepository)
	appHandler := application.NewHandler(applicationService)

	server := server.Server(appHandler)
	server.Run()

	logs.Info("API running")
}
