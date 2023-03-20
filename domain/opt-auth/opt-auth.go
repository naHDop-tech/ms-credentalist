package opt_auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	db "github.com/naHDop-tech/ms-credentalist/db/sqlc"
	email_sender "github.com/naHDop-tech/ms-credentalist/service/email-sender"
	"github.com/naHDop-tech/ms-credentalist/utils"
	"github.com/naHDop-tech/ms-credentalist/utils/opt"
)

var (
	createAuthRecordError  = errors.New("create auth record failed")
	customerNotExistsError = errors.New("customer not exists")
	generateOptCodeError   = errors.New("opt code generator failed")
	sendOptCodeError       = errors.New("sending opt code was failed")
)

type OptAuthDomain struct {
	repository *db.Queries
	sender     email_sender.Sender
}

func NewOptAuthDomain(conn *sql.DB, confg utils.Config) *OptAuthDomain {
	sender := email_sender.NewSmtpSender(confg)
	return &OptAuthDomain{
		repository: db.New(conn),
		sender:     sender,
	}
}

func (d OptAuthDomain) SentOpt(ctx context.Context, customerId uuid.UUID) error {
	customer, err := d.repository.GetUserByCustomerName(ctx, customerId)
	if err != nil {
		return err
	}
	if customer.ID.UUID == uuid.Nil {
		return customerNotExistsError
	}

	optCode, err := opt.GenerateOTP(6)
	if err != nil {
		fmt.Println("ERR", err.Error())
		return generateOptCodeError
	}

	textBody, err := email_sender.GetOptBodyMessage(optCode, "Verify yourself")
	if err != nil {
		fmt.Println("ERR", err.Error())
		return generateOptCodeError
	}

	to := []string{customer.Email}
	from := "tech.engineer.jedi@gmail.com"
	err = d.sender.Sent(from, to, textBody)
	if err != nil {
		fmt.Println("ERR", err.Error())
		return sendOptCodeError
	}

	_, err = d.repository.CreateAuthRecord(ctx, db.CreateAuthRecordParams{
		ID:         uuid.New(),
		IsVerified: false,
		Opt:        optCode,
		Channel:    "email",
		CustomerID: customerId,
	})
	if err != nil {
		return createAuthRecordError
	}

	return nil
}
