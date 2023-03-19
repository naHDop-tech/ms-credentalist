package api

type verifyEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type okResponse struct {
	Status string `json:"status"`
}
