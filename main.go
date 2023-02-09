package main

import (
	"database/sql"
	"fmt"
	"go-backend-template/api"
	router "go-backend-template/api/router"
	"go-backend-template/config"
	"go-backend-template/db"
	"go-backend-template/pkg/logging"
)

func init() {
	config.Setup()
	logging.Setup()
	db.Setup()
}

func main() {
	runGinServer(config.Cfg.Server.Address, db.SqlDB)
}

func runGinServer(address string, sqlDB *sql.DB) {
	handlers, err := Wire(sqlDB)
	if err != nil {
		panic(fmt.Errorf("cannot wire handlers: %v", err))

	}
	server := api.NewServer(router.SetupHandlers(handlers))

	err = server.Start(address)
	if err != nil {
		logging.Log.Error("cannot start server: %v", err)
	}
}
