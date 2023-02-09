package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go-backend-template/config"
	"go-backend-template/pkg/logging"
	"time"
)

var SqlDB *sql.DB

func Setup() {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Cfg.Database.Engine,
		config.Cfg.Database.User,
		config.Cfg.Database.Password,
		config.Cfg.Database.Host,
		config.Cfg.Database.Port,
		config.Cfg.Database.Name)
	conn, err := sql.Open(config.Cfg.Database.Engine, connStr)
	conn.SetMaxIdleConns(config.Cfg.Database.MaxIdle)
	conn.SetMaxOpenConns(config.Cfg.Database.MaxOpen)
	conn.SetConnMaxLifetime(time.Hour)

	if err != nil {
		logging.Log.Fatal("Failed to connect to the Database")
	}

	SqlDB = conn

	logging.Log.Info("🚀 Connected Successfully to the Database")
	runMigrations(connStr)
}

func runMigrations(connStr string) {
	migration, err := migrate.New(config.Cfg.Database.MigrationPath, connStr)
	if err != nil {
		logging.Log.Fatal("Failed to create migration instance: ", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		logging.Log.Fatal("Failed to run migration up: ", err)
	}

	logging.Log.Info("🚀 Migrations ran successfully")
}
