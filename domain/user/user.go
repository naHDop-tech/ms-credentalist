package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	db "github.com/naHDop-tech/ms-credentalist/db/sqlc"
	"github.com/naHDop-tech/ms-credentalist/utils"
)

var (
	beginTxError        = errors.New("begin transaction failed")
	createUserError     = errors.New("create user failed")
	createCustomerError = errors.New("create customer failed")
	customerNotFound    = errors.New("customer not found")
)

type UserDomain struct {
	repository *db.Queries
	conn       *sql.DB
}

func NewUserDomain(conn *sql.DB) *UserDomain {
	return &UserDomain{
		repository: db.New(conn),
		conn:       conn,
	}
}

func (u *UserDomain) CreateUser(ctx context.Context, dto CreateUserDto) (*uuid.UUID, error) {
	tx, err := u.conn.Begin()
	if err != nil {
		return nil, beginTxError
	}
	defer tx.Rollback()
	qtx := u.repository.WithTx(tx)

	userId, err := qtx.CreateUser(ctx, db.CreateUserParams{
		ID:       uuid.New(),
		Email:    dto.Email,
		UserName: dto.UserName,
	})
	if err != nil {
		return nil, createUserError
	}

	pwdHash, err := utils.HashAndSalt([]byte(dto.Email))
	if err != nil {
		return nil, err
	}
	customerId, err := qtx.CreateCustomer(ctx, db.CreateCustomerParams{
		ID:       uuid.New(),
		Password: pwdHash,
		UserName: dto.UserName,
		UserID:   userId,
	})
	if err != nil {
		return nil, createCustomerError
	}

	tx.Commit()
	return &customerId, nil
}

func (u UserDomain) GetCustomerByName(ctx context.Context, name string) (*db.Customer, error) {
	customer, err := u.repository.GetCustomerByUserName(ctx, name)
	if err != nil {
		return nil, customerNotFound
	}

	return &customer, nil
}
