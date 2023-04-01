package api

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/naHDop-tech/ms-credentalist/cmd/middleware"
	"github.com/naHDop-tech/ms-credentalist/utils/responser"
	"github.com/naHDop-tech/ms-credentalist/utils/token"
)

func (s *Server) customerById(ctx *gin.Context) {
	var response responser.Response
	var req customerIdRequestParam
	var err error
	err = ctx.ShouldBindUri(&req)
	if err != nil {
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}

	val := ctx.MustGet(middleware.AuthPayloadKey).(*token.Payload)
	if val.CustomerId != *req.CustomerId {
		err = errors.New("you do not have access to this user")
		response = s.responser.New(nil, err, responser.UNAUTHORIZED)
		ctx.JSON(response.Status, response)
		return
	}

	parsedCustomerId, err := uuid.Parse(*req.CustomerId)
	if err != nil {
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}

	customer, err := s.userDomain.GetCustomerById(ctx, parsedCustomerId)
	if err != nil {
		if err == sql.ErrNoRows {
			response = s.responser.New(nil, err, responser.NOT_FOUND)
			ctx.JSON(response.Status, response)
			return
		}
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}

	response = s.responser.New(customer, err, responser.OK)
	ctx.JSON(response.Status, response)
	return
}
