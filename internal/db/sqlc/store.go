package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

func (s *Store) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil /*&sql.TxOptions{}*/)
	if err != nil {
		return err
	}

	err = fn(New(tx))
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
