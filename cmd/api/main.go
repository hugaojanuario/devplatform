package main

import (
	"log"

	"github.com/hugaojanuario/devplatform/internal/application"
	"github.com/hugaojanuario/devplatform/internal/k8s"
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

	if _, err := k8s.NewClient(); err != nil {
		log.Fatalf("erro ao criar client k8s: %v", err)
	}

	applicationRepository := application.NewRepository(db)
	applicationService := application.NewService(applicationRepository)
	appHandler := application.NewHandler(applicationService)

	server := server.Server(appHandler)
	server.Run()

	logs.Info("API running")
}
