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
	createAuthRecordError      = errors.New("create auth record failed")
	customerNotExistsError     = errors.New("customer not exists")
	recordNotExistsError       = errors.New("record not exists")
	recordAlreadyVerifiedError = errors.New("record already verified")
	notValidOtpError           = errors.New("otp code not valid")
	generateOptCodeError       = errors.New("opt code generator failed")
	sendOptCodeError           = errors.New("sending opt code was failed")
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

func (o *OptAuthDomain) SendOpt(ctx context.Context, customerId uuid.UUID) error {
	customer, err := o.repository.GetUserByCustomerId(ctx, customerId)
	if err != nil {
		return err
	}
	if customer.ID.UUID == uuid.Nil {
		return customerNotExistsError
	}

	otpCode, err := opt.GenerateOTP(6)
	if err != nil {
		fmt.Println("ERR", err.Error())
		return generateOptCodeError
	}

	textBody, err := email_sender.GetOtpBodyMessage(otpCode, "Verify yourself")
	if err != nil {
		fmt.Println("ERR", err.Error())
		return generateOptCodeError
	}

	to := []string{customer.Email}
	from := "tech.engineer.jedi@gmail.com"
	err = o.sender.Sent(from, to, textBody)
	if err != nil {
		fmt.Println("ERR", err.Error())
		return sendOptCodeError
	}

	_, err = o.repository.CreateAuthRecord(ctx, db.CreateAuthRecordParams{
		ID:         uuid.New(),
		IsVerified: false,
		Otp:        otpCode,
		Channel:    "email",
		CustomerID: customerId,
	})
	if err != nil {
		return createAuthRecordError
	}

	return nil
}

func (o *OptAuthDomain) VerifyCustomerOtpCode(ctx context.Context, dto VerifyCustomerOtpDto) error {
	customer, err := o.repository.GetCustomerByUserName(ctx, dto.UserName)
	if err != nil {
		return customerNotExistsError
	}

	record, err := o.repository.GetLastRecord(ctx, customer.ID)
	if err != nil {
		return recordNotExistsError
	}
	if record.IsVerified {
		return recordAlreadyVerifiedError
	}
	if record.Otp != dto.Otp {
		return notValidOtpError
	}

	err = o.repository.VerifyCustomerOpt(ctx, db.VerifyCustomerOptParams{
		IsVerified: true,
		CustomerID: customer.ID,
	})

	return err
}
