package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-backend-template/repository/sqlc"
)

// Store provides all functions to execute database queries and transactions
type Store interface {
	sqlc.Querier
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	db *sql.DB
	*sqlc.Queries
}

// NewStore creates a new store
func NewStore(sqlDB *sql.DB) Store {
	return &SQLStore{
		db:      sqlDB,
		Queries: sqlc.New(sqlDB),
	}
}

// ExecTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := sqlc.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
