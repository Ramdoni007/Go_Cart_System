-- name: CreateCart :one
INSERT INTO cart (id,
name_product, quantity)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetCart :one
SELECT *
FROM cart
WHERE id = $1 LIMIT 1;

-- name: UpdateCartForQuantity :one
UPDATE cart
SET quantity = $2
WHERE id = $1 RETURNING *;

-- name: DeleteCart :exec
DELETE
FROM cart
WHERE id = $1;

-- name: GetQuantity :one
SELECT quantity
FROM cart
WHERE id = $1 LIMIT 1;

-- name: GetQuantityForUpdate :one
SELECT quantity
FROM cart
WHERE id = $1;
