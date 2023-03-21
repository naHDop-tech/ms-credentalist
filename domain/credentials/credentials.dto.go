package credentials

import "github.com/google/uuid"

type CreateCredentialDto struct {
	Title           string
	LoginName       string
	Secret          string
	Description     string
	ShowImmediately bool
	SendToEmail     bool
	SendToPhone     bool
	CustomerId      uuid.UUID
}

type UpdateCredentialDto struct {
	CreateCredentialDto
	CredentialId uuid.UUID
}
