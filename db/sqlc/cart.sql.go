// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: cart.sql

package db

import (
	"context"
)

const createCart = `-- name: CreateCart :one
INSERT INTO cart (id,
                  name_product, quantity)
VALUES ($1, $2, $3) RETURNING id, name_product, quantity
`

type CreateCartParams struct {
	ID          int64  `json:"id"`
	NameProduct string `json:"name_product"`
	Quantity    int64  `json:"quantity"`
}

func (q *Queries) CreateCart(ctx context.Context, arg CreateCartParams) (Cart, error) {
	row := q.db.QueryRowContext(ctx, createCart, arg.ID, arg.NameProduct, arg.Quantity)
	var i Cart
	err := row.Scan(&i.ID, &i.NameProduct, &i.Quantity)
	return i, err
}

const deleteCart = `-- name: DeleteCart :exec
DELETE
FROM cart
WHERE id = $1
`

func (q *Queries) DeleteCart(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCart, id)
	return err
}

const getCart = `-- name: GetCart :one
SELECT id, name_product, quantity
FROM cart
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCart(ctx context.Context, id int64) (Cart, error) {
	row := q.db.QueryRowContext(ctx, getCart, id)
	var i Cart
	err := row.Scan(&i.ID, &i.NameProduct, &i.Quantity)
	return i, err
}

const updateCart = `-- name: UpdateCart :one
UPDATE cart
SET name_product = $2,
    quantity     = $3
WHERE id = $1 RETURNING id, name_product, quantity
`

type UpdateCartParams struct {
	ID          int64  `json:"id"`
	NameProduct string `json:"name_product"`
	Quantity    int64  `json:"quantity"`
}

func (q *Queries) UpdateCart(ctx context.Context, arg UpdateCartParams) (Cart, error) {
	row := q.db.QueryRowContext(ctx, updateCart, arg.ID, arg.NameProduct, arg.Quantity)
	var i Cart
	err := row.Scan(&i.ID, &i.NameProduct, &i.Quantity)
	return i, err
}
