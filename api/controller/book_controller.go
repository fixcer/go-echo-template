package api

import (
	"github.com/gin-gonic/gin"
	"go-backend-template/api"
	"go-backend-template/mapper"
	"net/http"
)

func (c HandlerImpl) GetBooks(ctx *gin.Context, params api.GetBooksParams) {
	books := c.bookUseCase.FindAll()
	ctx.JSON(http.StatusOK, mapper.ToBookResponses(books))
}

func (c HandlerImpl) CreateBook(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (c HandlerImpl) GetBook(ctx *gin.Context, bookId string) {
	//TODO implement me
	panic("implement me")
}
