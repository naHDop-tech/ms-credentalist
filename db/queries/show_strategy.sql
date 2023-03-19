-- name: CreateShowStrategy :one
INSERT INTO show_strategies ("id", "show_immediately", "send_to_email", "send_to_phone", "credential_id")
VALUES ($1, $2, $3, $4, $5) RETURNING "id";

-- name: UpdateShowStrategy :exec
UPDATE show_strategies
SET show_immediately = $1, send_to_email = $2, send_to_phone = $3, updated_at = $4
WHERE id = $5;

-- name: GetShowStrategy :one
SELECT * FROM show_strategies
WHERE credential_id = $1 LIMIT 1;