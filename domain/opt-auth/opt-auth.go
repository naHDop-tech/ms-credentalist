package opt_auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	db "github.com/naHDop-tech/ms-credentalist/db/sqlc"
)

var (
	beginTxError           = errors.New("begin transaction failed")
	createAuthRecordError  = errors.New("create customer failed")
	customerNotExistsError = errors.New("customer not exists")
)

type OptAuthDomain struct {
	repository *db.Queries
}

func NewOptAuthDomain(conn *sql.DB) *OptAuthDomain {
	return &OptAuthDomain{
		repository: db.New(conn),
	}
}

func (d OptAuthDomain) SentOpt(ctx context.Context, customerId uuid.UUID) error {
	customer, err := d.repository.GetCustomerById(ctx, customerId)
	if err != nil {
		return err
	}
	if customer.ID == uuid.Nil {
		return customerNotExistsError
	}
	_, err = d.repository.CreateAuthRecord(ctx, db.CreateAuthRecordParams{
		ID:         uuid.New(),
		IsVerified: false,
		// TODO: get opt from service
		Opt:        "735162",
		Channel:    "email",
		CustomerID: customerId,
	})
	if err != nil {
		return createAuthRecordError
	}

	return nil
}
