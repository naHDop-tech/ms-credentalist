package opt_auth

type CreateOptAuthRecord struct {
	Email    string
	UserName string
}

type VerifyCustomerOtpDto struct {
	Otp      string
	UserName string
}
