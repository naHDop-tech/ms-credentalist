-- name: CreateAuthRecord :one
INSERT INTO customer_auth ("id", "is_verified", "opt", "channel", "customer_id")
VALUES ($1, $2, $3, $4, $5) RETURNING "id";

-- name: GetAuthRecordByCustomerId :one
SELECT * FROM customer_auth
WHERE customer_id = $1 ORDER BY created_at ASC LIMIT 1;

-- name: GetAuthRecordHistory :many
SELECT * FROM customer_auth
WHERE customer_id = $1 ORDER BY created_at ASC;

-- name: GetLastRecord :one
SELECT * FROM customer_auth
WHERE customer_id = $1 ORDER BY created_at ASC LIMIT 1;