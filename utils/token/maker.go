package token

import (
	"time"
)

type UserPayload struct {
	CustomerId string
	UserName   string
}

type Maker interface {
	CreateToken(userPayload UserPayload, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
