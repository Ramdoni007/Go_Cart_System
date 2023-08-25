-- name: CreateCart :one
INSERT INTO cart (id,
                  name_product, quantity)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetCart :one
SELECT
FROM cart
WHERE id = $1 LIMIT 1;

-- name: UpdateCart :one
UPDATE cart
SET name_product = $2,
    quantity     = $3
WHERE id = $1 RETURNING *;

-- name: DeleteCart :exec
DELETE
FROM cart
WHERE id = $1;
