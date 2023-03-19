// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: credentials.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const userCredentials = `-- name: UserCredentials :many
SELECT
    c.id as "credential_id",
    c.title as "title",
    c.login_name as "login_name",
    c.secret as "secret",
    c.description as "description",
    ss.show_immediately as "show_immediately",
    ss.send_to_email as "send_to_email",
    ss.send_to_phone as "send_to_phone"
FROM credentials c
JOIN show_strategies ss ON ss.credential_id = c.id
WHERE c.customer_id = $1
`

type UserCredentialsRow struct {
	CredentialID    uuid.UUID `json:"credential_id"`
	Title           string    `json:"title"`
	LoginName       string    `json:"login_name"`
	Secret          string    `json:"secret"`
	Description     string    `json:"description"`
	ShowImmediately bool      `json:"show_immediately"`
	SendToEmail     bool      `json:"send_to_email"`
	SendToPhone     bool      `json:"send_to_phone"`
}

func (q *Queries) UserCredentials(ctx context.Context, customerID uuid.UUID) ([]UserCredentialsRow, error) {
	rows, err := q.db.QueryContext(ctx, userCredentials, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserCredentialsRow{}
	for rows.Next() {
		var i UserCredentialsRow
		if err := rows.Scan(
			&i.CredentialID,
			&i.Title,
			&i.LoginName,
			&i.Secret,
			&i.Description,
			&i.ShowImmediately,
			&i.SendToEmail,
			&i.SendToPhone,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
