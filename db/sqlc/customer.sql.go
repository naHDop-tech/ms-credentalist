// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: customer.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createCustomer = `-- name: CreateCustomer :one
INSERT INTO customers ("id", "password", "user_name", "user_id")
VALUES ($1, $2, $3, $4) RETURNING "id"
`

type CreateCustomerParams struct {
	ID       uuid.UUID `json:"id"`
	Password string    `json:"password"`
	UserName string    `json:"user_name"`
	UserID   uuid.UUID `json:"user_id"`
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, createCustomer,
		arg.ID,
		arg.Password,
		arg.UserName,
		arg.UserID,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getCustomerById = `-- name: GetCustomerById :one
SELECT id, password, user_name, user_id, created_at, updated_at FROM customers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCustomerById(ctx context.Context, id uuid.UUID) (Customer, error) {
	row := q.db.QueryRowContext(ctx, getCustomerById, id)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Password,
		&i.UserName,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCustomerByUserName = `-- name: GetCustomerByUserName :one
SELECT id, password, user_name, user_id, created_at, updated_at FROM customers
WHERE user_name = $1 LIMIT 1
`

func (q *Queries) GetCustomerByUserName(ctx context.Context, userName string) (Customer, error) {
	row := q.db.QueryRowContext(ctx, getCustomerByUserName, userName)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Password,
		&i.UserName,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByCustomerId = `-- name: GetUserByCustomerId :one
SELECT
    u.email,
    c.id,
    c.user_name,
    c.user_id,
    c.created_at,
    c.updated_at
FROM users u
LEFT JOIN customers c ON c.user_id = u.id
WHERE c.id = $1 LIMIT 1
`

type GetUserByCustomerIdRow struct {
	Email     string         `json:"email"`
	ID        uuid.NullUUID  `json:"id"`
	UserName  sql.NullString `json:"user_name"`
	UserID    uuid.NullUUID  `json:"user_id"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
}

func (q *Queries) GetUserByCustomerId(ctx context.Context, id uuid.UUID) (GetUserByCustomerIdRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByCustomerId, id)
	var i GetUserByCustomerIdRow
	err := row.Scan(
		&i.Email,
		&i.ID,
		&i.UserName,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePassword = `-- name: UpdatePassword :exec
UPDATE customers
SET password = $1, updated_at = $2
WHERE id = $3
`

type UpdatePasswordParams struct {
	Password  string       `json:"password"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	ID        uuid.UUID    `json:"id"`
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error {
	_, err := q.db.ExecContext(ctx, updatePassword, arg.Password, arg.UpdatedAt, arg.ID)
	return err
}

const updateUserName = `-- name: UpdateUserName :exec
UPDATE customers
SET user_name = $1, updated_at = $2
WHERE id = $3
`

type UpdateUserNameParams struct {
	UserName  string       `json:"user_name"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	ID        uuid.UUID    `json:"id"`
}

func (q *Queries) UpdateUserName(ctx context.Context, arg UpdateUserNameParams) error {
	_, err := q.db.ExecContext(ctx, updateUserName, arg.UserName, arg.UpdatedAt, arg.ID)
	return err
}
