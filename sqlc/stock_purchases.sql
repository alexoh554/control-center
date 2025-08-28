-- name: CreateStockPurchase :one
INSERT INTO stock_purchases(
    symbol,
    price_cents,
    quantity,
    purchased_at,
    total_price_cents,
    broker
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: DeleteStockPurchase :exec
DELETE FROM stock_purchases
WHERE id=$1;