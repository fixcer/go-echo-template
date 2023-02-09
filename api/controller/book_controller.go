package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"go-backend-template/api"
	"go-backend-template/pkg/logging"
	"net/http"
)

func (c HandlerImpl) GetBooks(ctx *gin.Context, params api.GetBooksParams) {
	books := c.bookUseCase.FindAll()
	var bookResponses *[]api.Book
	err := mapstructure.Decode(books, &bookResponses)
	if err != nil {
		logging.Log.Error("Error while decoding books: ", err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, bookResponses)
}

func (c HandlerImpl) CreateBook(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c HandlerImpl) GetBook(ctx *gin.Context, bookId string) {
	//TODO implement me
	panic("implement me")
}
