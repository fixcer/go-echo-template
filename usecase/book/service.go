package book

import (
	"context"
	"go-backend-template/pkg/logging"
	"go-backend-template/repository"
	"go-backend-template/repository/sqlc"
	"sync"
)

var (
	serviceOne      sync.Once
	serviceInstance *service
)

type service struct {
	store repository.Store
}

// NewBookUseCase returns a new book usecase
func NewBookUseCase(store repository.Store) UseCase {
	serviceOne.Do(func() {
		serviceInstance = &service{
			store: store,
		}
	})

	return serviceInstance
}

// FindAll returns all books
func (s *service) FindAll() []sqlc.Book {
	books, err := s.store.ListBooks(context.Background(), sqlc.ListBooksParams{
		Limit:  10,
		Offset: 0,
	})

	if err != nil {
		logging.Log.Error("cannot get books: %v", err)
		return []sqlc.Book{}
	}

	return books
}
