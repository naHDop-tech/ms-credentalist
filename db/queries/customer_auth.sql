-- name: CreateAuthSettings :one
INSERT INTO customer_auth ("id", "is_verified", "opt", "channel", "customer_id")
VALUES ($1, $2, $3, $4, $5) RETURNING "id";

-- name: GetAuthSettingsByCustomerId :one
SELECT * FROM customer_auth
WHERE customer_id = $1 ORDER BY created_at ASC LIMIT 1;

-- name: GetAuthSettingsHistory :many
SELECT * FROM customer_auth
WHERE customer_id = $1 ORDER BY created_at ASC;

-- name: GetLastNotVerifiedRecord :one
SELECT * FROM customer_auth
WHERE customer_id = $1 AND is_verified = false ORDER BY created_at ASC LIMIT 1;