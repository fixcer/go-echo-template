package api

import "go-backend-template/usecase/book"

// HandlerImpl is the implementation of the ServerInterface.
type HandlerImpl struct {
	bookUseCase book.UseCase
}

// NewHandlerImpl creates a new HandlerImpl.
func NewHandlerImpl(bookUseCase book.UseCase) *HandlerImpl {
	return &HandlerImpl{
		bookUseCase: bookUseCase,
	}
}
