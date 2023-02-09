package main

import (
	"go-backend-template/api"
	router "go-backend-template/api/router"
	"go-backend-template/config"
	"go-backend-template/db"
	"go-backend-template/pkg/logging"
	"go-backend-template/repository"
)

func init() {
	config.Setup()
	logging.Setup()
	db.Setup()
}

func main() {
	runGinServer(config.Cfg.Server.Address, repository.NewStore(db.SqlDB))
}

func runGinServer(address string, store repository.Store) {
	server := api.NewServer(router.SetupHandlers(store))

	err := server.Start(address)
	if err != nil {
		logging.Log.Error("cannot start server: %v", err)
	}
}
