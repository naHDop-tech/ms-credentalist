// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users ("id", "email", "user_name")
VALUES ($1, $2, $3) RETURNING "id"
`

type CreateUserParams struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	UserName string    `json:"user_name"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.ID, arg.Email, arg.UserName)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}
