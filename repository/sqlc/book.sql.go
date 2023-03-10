// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: book.sql

package sqlc

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createBook = `-- name: CreateBook :one
insert into books (code, title, publisher, public_date) values ($1, $2, $3, $4) returning book_id, code, title, publisher, public_date, version, created_at, updated_at
`

type CreateBookParams struct {
	Code       string    `json:"code"`
	Title      string    `json:"title"`
	Publisher  string    `json:"publisher"`
	PublicDate time.Time `json:"public_date"`
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, createBook,
		arg.Code,
		arg.Title,
		arg.Publisher,
		arg.PublicDate,
	)
	var i Book
	err := row.Scan(
		&i.BookID,
		&i.Code,
		&i.Title,
		&i.Publisher,
		&i.PublicDate,
		&i.Version,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getBook = `-- name: GetBook :one
select book_id, code, title, publisher, public_date, version, created_at, updated_at from books
where book_id = $1 limit 1
`

func (q *Queries) GetBook(ctx context.Context, bookID uuid.UUID) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBook, bookID)
	var i Book
	err := row.Scan(
		&i.BookID,
		&i.Code,
		&i.Title,
		&i.Publisher,
		&i.PublicDate,
		&i.Version,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listBooks = `-- name: ListBooks :many
select book_id, code, title, publisher, public_date, version, created_at, updated_at from books
order by created_at
limit $1
offset $2
`

type ListBooksParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListBooks(ctx context.Context, arg ListBooksParams) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, listBooks, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Book{}
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.BookID,
			&i.Code,
			&i.Title,
			&i.Publisher,
			&i.PublicDate,
			&i.Version,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
