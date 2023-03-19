package domain

import (
	"context"

	"github.com/google/uuid"
)

type Domain interface {
	GetByCustomerName(context.Context, string) (any, error)
	GetById(context.Context, uuid.UUID) (any, error)
	Update(context.Context, any) (any, error)
	Create(context.Context, any) (any, error)
	GetListByCustomerId(context.Context, uuid.UUID) (any, error)
}
