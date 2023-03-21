package api

type getUserByIdRequest struct {
	CustomerId *string `uri:"customer_id" binding:"omitempty,uuid"`
}
