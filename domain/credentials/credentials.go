package credentials

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	db "github.com/naHDop-tech/ms-credentalist/db/sqlc"
)

var (
	customerNotExistsError    = errors.New("customer not exists")
	credentialsNotExistsError = errors.New("credentials not exists")
	beginTxError              = errors.New("begin transaction failed")
	credentialNotFoundError   = errors.New("credential not found")
	showStrategyNotFoundError = errors.New("show strategy not found")
)

type CredentialsDomain struct {
	repository *db.Queries
	conn       *sql.DB
}

func NewCredentialsDomain(conn *sql.DB) *CredentialsDomain {
	return &CredentialsDomain{
		repository: db.New(conn),
		conn:       conn,
	}
}

func (c *CredentialsDomain) GetCredentials(ctx context.Context, customerId uuid.UUID) (*[]db.CustomerCredentialsRow, error) {
	customer, err := c.repository.GetCustomerById(ctx, customerId)
	if err != nil {
		return nil, err
	}
	if customer.ID == uuid.Nil {
		return nil, customerNotExistsError
	}

	credentials, err := c.repository.CustomerCredentials(ctx, customer.ID)
	if err != nil {
		return nil, err
	}
	if len(credentials) == 0 {
		return nil, credentialsNotExistsError
	}

	return &credentials, nil
}

func (c *CredentialsDomain) CreateCredential(ctx context.Context, dto CreateCredentialDto) error {
	tx, err := c.conn.Begin()
	if err != nil {
		return beginTxError
	}
	defer tx.Rollback()
	qtx := c.repository.WithTx(tx)

	customer, err := qtx.GetCustomerById(ctx, dto.CustomerId)
	if err != nil {
		return err
	}
	if customer.ID == uuid.Nil {
		return customerNotExistsError
	}

	credentialId, err := qtx.AddCredential(ctx, db.AddCredentialParams{
		ID:          uuid.New(),
		Title:       dto.Title,
		LoginName:   dto.LoginName,
		Secret:      dto.Secret,
		Description: dto.Description,
		CustomerID:  dto.CustomerId,
	})
	if err != nil {
		return err
	}

	_, err = qtx.CreateShowStrategy(ctx, db.CreateShowStrategyParams{
		ID:              uuid.New(),
		ShowImmediately: dto.ShowImmediately,
		SendToEmail:     dto.SendToEmail,
		SendToPhone:     dto.SendToPhone,
		CredentialID:    credentialId,
	})
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func (c *CredentialsDomain) UpdateCredential(ctx context.Context, dto UpdateCredentialDto) error {
	tx, err := c.conn.Begin()
	if err != nil {
		return beginTxError
	}
	defer tx.Rollback()
	qtx := c.repository.WithTx(tx)

	customer, err := qtx.GetCustomerById(ctx, dto.CustomerId)
	if err != nil {
		return err
	}
	if customer.ID == uuid.Nil {
		return customerNotExistsError
	}

	credential, err := qtx.GetCredential(ctx, dto.CredentialId)
	if err != nil {
		return err
	}
	if credential.ID == uuid.Nil {
		return credentialNotFoundError
	}

	showStrategy, err := qtx.GetShowStrategy(ctx, credential.ID)
	if err != nil {
		return err
	}
	if showStrategy.ID == uuid.Nil {
		return showStrategyNotFoundError
	}

	err = qtx.UpdateShowStrategy(ctx, db.UpdateShowStrategyParams{
		ShowImmediately: dto.ShowImmediately,
		SendToEmail:     dto.SendToEmail,
		SendToPhone:     dto.SendToPhone,
		UpdatedAt:       sql.NullTime{Valid: true, Time: time.Now()},
		ID:              showStrategy.ID,
	})
	if err != nil {
		return err
	}

	err = qtx.UpdateCredential(ctx, db.UpdateCredentialParams{
		Title:       dto.Title,
		LoginName:   dto.LoginName,
		Secret:      dto.Secret,
		Description: dto.Description,
		UpdatedAt:   sql.NullTime{Valid: true, Time: time.Now()},
		ID:          credential.ID,
	})

	tx.Commit()
	return nil
}
