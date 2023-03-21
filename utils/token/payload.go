package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToke  = errors.New("invalid toke")
)

type Payload struct {
	ID         uuid.UUID `json:"id"`
	CustomerId string    `json:"user_id"`
	UserName   string    `json:"user_name"`
	IssuedAt   time.Time `json:"issued_at"`
	ExpiredAt  time.Time `json:"expired_at"`
}

func CreatePayload(userPayload UserPayload, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:         tokenId,
		CustomerId: userPayload.CustomerId,
		UserName:   userPayload.UserName,
		IssuedAt:   time.Now(),
		ExpiredAt:  time.Now().Add(duration),
	}

	return payload, nil
}
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}

	return nil
}
