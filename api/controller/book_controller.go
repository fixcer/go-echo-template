package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"go-backend-template/api"
	"go-backend-template/usecase/book"
	"net/http"
)

type BookController struct {
	BookUsecase book.UseCase
}

func (c ServerImpl) GetBooks(ctx *gin.Context, params api.GetBooksParams) {
	books := c.BookController.BookUsecase.FindAll()
	var bookResponses *[]api.Book
	err := mapstructure.Decode(books, &bookResponses)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, bookResponses)
}

func (c ServerImpl) CreateBook(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c ServerImpl) GetBook(ctx *gin.Context, bookId string) {
	//TODO implement me
	panic("implement me")
}
