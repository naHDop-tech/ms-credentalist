-- name: CreateAuthRecord :one
INSERT INTO customer_auth ("id", "is_verified", "otp", "channel", "customer_id")
VALUES ($1, $2, $3, $4, $5) RETURNING "id";

-- name: VerifyCustomerOpt :exec
UPDATE customer_auth
SET is_verified = $1 WHERE customer_id = $2;

-- name: GetAuthRecordByCustomerId :one
SELECT * FROM customer_auth
WHERE customer_id = $1 ORDER BY created_at DESC LIMIT 1;

-- name: GetAuthRecordHistory :many
SELECT * FROM customer_auth
WHERE customer_id = $1 ORDER BY created_at DESC;

-- name: GetLastRecord :one
SELECT * FROM customer_auth
WHERE customer_id = $1 ORDER BY created_at DESC LIMIT 1;