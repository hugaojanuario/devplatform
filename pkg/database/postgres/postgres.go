package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/hugaojanuario/devplatform/pkg/config"
	_ "github.com/lib/pq"
)

func Conn(cfg *config.Config) (*sql.DB, error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DbName, cfg.Postgres.Sslmode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error establishing a connection to the database: %w", err)
	}

	return db, nil
}
