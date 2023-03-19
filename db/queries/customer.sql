-- name: CreateCustomer :one
INSERT INTO customers ("id", "password", "user_name", "user_id")
VALUES ($1, $2, $3, $4) RETURNING "id";

-- name: GetCustomerByUserName :one
SELECT * FROM customers
WHERE user_name = $1 LIMIT 1;

-- name: GetCustomerById :one
SELECT * FROM customers
WHERE id = $1 LIMIT 1;

-- name: UpdateUserName :exec
UPDATE customers
SET user_name = $1, updated_at = $2
WHERE id = $3;

-- name: UpdatePassword :exec
UPDATE customers
SET password = $1, updated_at = $2
WHERE id = $3;