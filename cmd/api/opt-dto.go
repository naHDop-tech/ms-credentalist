package api

type verifyEmailRequest struct {
	Email    string `json:"email" binding:"required,email"`
	UserName string `json:"user_name" binding:"required,min=5,max=30"`
}

type verifyOtpRequest struct {
	Otp      string `json:"otp" binding:"required,min=6,max=6"`
	UserName string `json:"user_name" binding:"required,min=5,max=30"`
}

type okResponse struct {
	Status string `json:"status"`
}

type tokenResponse struct {
	Token string `json:"token"`
}
