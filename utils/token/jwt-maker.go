package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const minSecretSize = 32

type JWTMaker struct {
	secret string
}

func NewJWTMaker(secret string) (Maker, error) {
	if len(secret) < minSecretSize {
		return nil, errors.New("to short secret size")
	}

	return &JWTMaker{secret}, nil
}

func (tm *JWTMaker) CreateToken(userPayload UserPayload, duration time.Duration) (string, error) {
	payload, err := CreatePayload(userPayload, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(tm.secret))
}

func (tm *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToke
		}

		return []byte(tm.secret), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToke
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToke
	}

	return payload, nil
}
