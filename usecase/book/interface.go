package book

import "go-backend-template/repository/sqlc"

type UseCase interface {
	FindAll() []sqlc.Book
}
