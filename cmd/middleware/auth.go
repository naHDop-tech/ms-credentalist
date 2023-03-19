package middleware

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/naHDop-tech/ms-credentalist/utils/responser"
	"github.com/naHDop-tech/ms-credentalist/utils/token"
)

const (
	authHeader     = "authorization"
	authType       = "bearer"
	authPayloadKey = "auth_payload"
)

func AuthMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authHeader)
		r := responser.NewResponser()
		var response responser.Response
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header has not provided")
			response = r.New(nil, err, responser.UNAUTHORIZED)
			ctx.AbortWithStatusJSON(response.Status, response)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("authorization header invalid format")
			response = r.New(nil, err, responser.UNAUTHORIZED)
			ctx.AbortWithStatusJSON(response.Status, response)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authType {
			err := errors.New("authorization header invalid type")
			response = r.New(nil, err, responser.UNAUTHORIZED)
			ctx.AbortWithStatusJSON(response.Status, response)
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			response = r.New(nil, err, responser.UNAUTHORIZED)
			ctx.AbortWithStatusJSON(response.Status, response)
			return
		}

		ctx.Set(authPayloadKey, payload)
		ctx.Next()
	}
}
