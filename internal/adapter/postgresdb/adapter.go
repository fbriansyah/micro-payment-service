package postgresdb

import (
	"database/sql"
)

type DatabaseAdapter interface {
	Querier
}

type DatabaseStore struct {
	db *sql.DB
	*Queries
}

func NewDatabaseAdapter(db *sql.DB) DatabaseAdapter {
	return &DatabaseStore{
		db:      db,
		Queries: New(db),
	}
}

// func (store *DatabaseStore) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := store.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	q := New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
// 		}
// 		return err
// 	}

// 	return tx.Commit()
// }
