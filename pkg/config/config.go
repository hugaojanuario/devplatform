package config

import (
	"os"
	"path/filepath"

	"github.com/hugaojanuario/devplatform/internal/logger"
	"github.com/joho/godotenv"
)

type PostgresConfig struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
	Sslmode  string
}

type ApiConfig struct {
	Url  string
	Port string
}

type Config struct {
	Postgres PostgresConfig
	Api      ApiConfig
}

func LoadEnvFile() *Config {
	logs := logger.Logger()
	found := false

	for i := 0; i < 10; i++ {
		envPath := ".env"
		for j := 0; j < i; j++ {
			envPath = filepath.Join("..", envPath)
		}
		err := godotenv.Load(envPath)
		if err != nil {
			continue
		} else {
			found = true
			break
		}

	}
	if !found {
		logs.Info("no .env file found; relying on environment variables")
	}

	return &Config{
		Postgres: PostgresConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			DbName:   os.Getenv("POSTGRES_DB"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Sslmode:  os.Getenv("POSTGRES_SSLMODE"),
		},
		Api: ApiConfig{
			Url:  os.Getenv("API_URL"),
			Port: os.Getenv("API_PORT"),
		},
	}
}
