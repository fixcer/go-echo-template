//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"
	controller "go-backend-template/api/controller"
	"go-backend-template/repository"
	"go-backend-template/usecase/book"
)

var HandlerSet = wire.NewSet(controller.NewHandlerImpl)
var UseCaseSet = wire.NewSet(book.NewBookUseCase)
var RepositorySet = wire.NewSet(repository.NewStore)

// Wire is a function that wires up the application.
func Wire(sqlDB *sql.DB) (*controller.HandlerImpl, error) {
	panic(wire.Build(HandlerSet, UseCaseSet, RepositorySet))
}
