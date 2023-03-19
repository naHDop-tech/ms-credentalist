// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Credential struct {
	ID          uuid.UUID    `json:"id"`
	Title       string       `json:"title"`
	LoginName   string       `json:"login_name"`
	Secret      string       `json:"secret"`
	Description string       `json:"description"`
	CustomerID  uuid.UUID    `json:"customer_id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   sql.NullTime `json:"updated_at"`
}

type Customer struct {
	ID        uuid.UUID    `json:"id"`
	Password  string       `json:"password"`
	UserName  string       `json:"user_name"`
	UserID    uuid.UUID    `json:"user_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type CustomerAuth struct {
	ID         uuid.UUID `json:"id"`
	CustomerID uuid.UUID `json:"customer_id"`
	IsVerified bool      `json:"is_verified"`
	Opt        string    `json:"opt"`
	Channel    string    `json:"channel"`
	CreatedAt  time.Time `json:"created_at"`
}

type ShowStrategy struct {
	ID              uuid.UUID    `json:"id"`
	ShowImmediately bool         `json:"show_immediately"`
	SendToEmail     bool         `json:"send_to_email"`
	SendToPhone     bool         `json:"send_to_phone"`
	CredentialID    uuid.UUID    `json:"credential_id"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
}

type User struct {
	ID        uuid.UUID    `json:"id"`
	Email     string       `json:"email"`
	UserName  string       `json:"user_name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
