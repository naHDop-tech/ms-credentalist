-- name: AddCredential :one
INSERT INTO credentials ("id", "title", "login_name", "secret", "description", "customer_id")
VALUES ($1, $2, $3, $4, $5, $6) RETURNING "id";

-- name: GetCredentialsByCustomerId :many
SELECT * FROM credentials
WHERE customer_id = $1;

-- name: UpdateCredential :exec
UPDATE credentials
SET title = $1, login_name = $2, secret = $3, description = $4, updated_at = $5
WHERE id = $6;