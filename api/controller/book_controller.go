package api

import (
	"github.com/labstack/echo/v4"
	"go-backend-template/api"
	"go-backend-template/mapper"
	"net/http"
)

func (c HandlerImpl) GetBooks(ctx echo.Context, params api.GetBooksParams) error {
	books := c.bookUseCase.FindAll()
	err := ctx.JSON(http.StatusOK, mapper.ToBookResponses(books))
	if err != nil {
		panic(err)
	}

	return nil
}

func (c HandlerImpl) CreateBook(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (c HandlerImpl) GetBook(ctx echo.Context, bookId string) error {
	//TODO implement me
	panic("implement me")
}
