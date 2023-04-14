package database

import (
	"embed"
	"fmt"

	"github.com/Wave-ETH-Global/wave-node/config"
	"github.com/google/logger"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrations embed.FS

// TODO refactor it all later and create abstract interface for storage
func ProvideDatabase(cfg *config.Config) *sqlx.DB {
	if cfg.DB.Type != "postgres" {
		logger.Fatal("The only supported database is PostgreSQL!")
	}

	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.DB.User, cfg.DB.Password, cfg.DB.Address, cfg.DB.Port, cfg.DB.DatabaseName))
	if err != nil {
		logger.Fatal(err)
	}

	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("postgres"); err != nil {
		logger.Fatal(err)
	}

	if err := goose.Up(db.DB, "migrations"); err != nil {
		logger.Fatal(err)
	}

	return db
}
