-- name: CreateUser :one
INSERT INTO users ("id", "email", "user_name")
VALUES ($1, $2, $3) RETURNING "id";