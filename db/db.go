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

var sqlDB *sql.DB

// Setup creates a new database connection and runs migrations
func Setup() {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Cfg.Database.Engine,
		config.Cfg.Database.User,
		config.Cfg.Database.Password,
		config.Cfg.Database.Host,
		config.Cfg.Database.Port,
		config.Cfg.Database.Name)

	var err error
	sqlDB, err = sql.Open(config.Cfg.Database.Engine, connStr)
	sqlDB.SetMaxIdleConns(config.Cfg.Database.MaxIdle)
	sqlDB.SetMaxOpenConns(config.Cfg.Database.MaxOpen)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		logging.Log.Fatal("Failed to connect to the Database")
	}

	logging.Log.Info("🚀 Connected Successfully to the Database")
	runMigrations(connStr)
}

// Instance returns the database connection
func Instance() *sql.DB {
	return sqlDB
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
