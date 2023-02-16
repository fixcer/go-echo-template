package mapper

import (
	"github.com/deepmap/oapi-codegen/pkg/types"
	"go-backend-template/api"
	"go-backend-template/repository/sqlc"
)

// ToBookResponses converts sqlc.Book to api.Book
func ToBookResponses(books []sqlc.Book) []api.Book {
	bookResponses := make([]api.Book, len(books))

	for idx := range books {
		book := books[idx]
		bookResponses[idx] = api.Book{
			BookId:    book.BookID,
			Code:      book.Code,
			Title:     &book.Title,
			Publisher: &book.Publisher,
			PublicDate: &types.Date{
				Time: book.PublicDate,
			},
		}
	}

	return bookResponses
}
