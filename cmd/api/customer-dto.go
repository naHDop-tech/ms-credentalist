package api

type customerIdRequestParam struct {
	CustomerId *string `uri:"customer_id" binding:"omitempty,uuid"`
}
