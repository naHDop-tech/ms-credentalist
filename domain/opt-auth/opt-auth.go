package opt_auth

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	db "github.com/naHDop-tech/ms-credentalist/db/sqlc"
	"github.com/naHDop-tech/ms-credentalist/utils"
)

var (
	beginTxError          = errors.New("begin transaction failed")
	createUserError       = errors.New("create user failed")
	createCustomerError   = errors.New("create customer failed")
	createAuthRecordError = errors.New("create customer failed")
)

type OptAuthDomain struct {
	repository *db.Queries
	conn       *sql.DB
}

func NewOptAuthDomain(conn *sql.DB) *OptAuthDomain {
	return &OptAuthDomain{
		repository: db.New(conn),
		conn:       conn,
	}
}

func (d OptAuthDomain) SentOpt(ctx context.Context, payload CreateOptAuthRecord) error {
	tx, err := d.conn.Begin()
	if err != nil {
		return beginTxError
	}
	defer tx.Rollback()
	qtx := d.repository.WithTx(tx)

	// TODO: check before create user

	userId, err := qtx.CreateUser(ctx, db.CreateUserParams{
		ID:       uuid.New(),
		Email:    payload.Email,
		UserName: payload.UserName,
	})
	if err != nil {
		return createUserError
	}

	pwdHash, err := utils.HashAndSalt([]byte(payload.Email))
	if err != nil {
		return err
	}
	customerId, err := qtx.CreateCustomer(ctx, db.CreateCustomerParams{
		ID:       uuid.New(),
		Password: pwdHash,
		UserName: payload.UserName,
		UserID:   userId,
	})
	if err != nil {
		return createCustomerError
	}

	_, err = qtx.CreateAuthRecord(ctx, db.CreateAuthRecordParams{
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

	return tx.Commit()
}
