-- name: CreateBook :one
insert into books (code, title, publisher, public_date) values ($1, $2, $3, $4) returning *;

-- name: GetBook :one
select * from books
where book_id = $1 limit 1;

-- name: ListBooks :many
select * from books
order by created_at
limit $1
offset $2;
