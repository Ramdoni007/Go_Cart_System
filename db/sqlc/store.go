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
		db:      db,
		Queries: New(db),
	}
}

//execTx to execute a function within a databases transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// BeginTx for Start Db Transaction
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)

	if err != nil {
		//RollBack For CallBack in Db Transaction.
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	//Commit For Apply Db Transaction
	return tx.Commit()
}

type QuantityTxParams struct {
	ID          int64  `json:"id"`
	NameProduct string `json:"name_product"`
	Quantity    int64  `json:"quantity"`
}

type QuantityTxResult struct {
	ID             Cart  `json:"id"`
	BeforeQuantity int64 `json:"before_quantity"`
	EditQuantity   int64 `json:"edit_quantity"`
}

func (store *Store) QuantityTx(ctx context.Context, arg QuantityTxParams) (QuantityTxResult, error) {
	var result QuantityTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.ID, err = q.CreateCart(ctx, CreateCartParams{
			ID:       arg.ID,
			Quantity: arg.Quantity,
		})
		if err != nil {
			return err
		}

		result.BeforeQuantity, err = q.GetQuantity(ctx, arg.Quantity)
		if err != nil {
			return err
		}

		result.EditQuantity, err = q.GetQuantityForUpdate(ctx, arg.Quantity)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
