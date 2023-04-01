package api

type createCredentialRequest struct {
	Title           string `json:"title"`
	LoginName       string `json:"login_name"`
	Secret          string `json:"secret"`
	Description     string `json:"description"`
	ShowImmediately bool   `json:"show_immediately"`
	SendToEmail     bool   `json:"send_to_email"`
	SendToPhone     bool   `json:"send_to_phone"`
}

type updateCredentialRequest struct {
	Title           string `json:"title"`
	LoginName       string `json:"login_name"`
	Secret          string `json:"secret"`
	Description     string `json:"description"`
	ShowImmediately bool   `json:"show_immediately"`
	SendToEmail     bool   `json:"send_to_email"`
	SendToPhone     bool   `json:"send_to_phone"`
}

type updateCredentialRequestParams struct {
	CustomerId   *string `uri:"customer_id" binding:"omitempty,uuid"`
	CredentialId *string `uri:"credential_id" binding:"omitempty,uuid"`
}
