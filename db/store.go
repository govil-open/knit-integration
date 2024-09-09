package db

import (
	"database/sql"
	db "knit-integration/db/sqlc"
)

// Store defines all functions to execute db queries and transactions.
type Store interface {
	db.Querier
}

// SQLStore provides all functions to execute SQL queries and transactions.
type SQLStore struct {
	db *sql.DB
	*db.Queries
}

func NewStore(sqldb *sql.DB) Store {
	return &SQLStore{
		db:      sqldb,
		Queries: db.New(sqldb),
	}
}

// // ExecTx executes a function within a database transaction
// func (store *SQLStore) execTx(ctx context.Context, fn func(*db.Queries) error) error {
// 	tx, err := store.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}
// 	q := db.New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
// 		}
// 		return err
// 	}
// 	return tx.Commit()
// }.
