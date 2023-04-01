package api

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/naHDop-tech/ms-credentalist/cmd/middleware"
	"github.com/naHDop-tech/ms-credentalist/domain/credentials"
	"github.com/naHDop-tech/ms-credentalist/utils/responser"
	"github.com/naHDop-tech/ms-credentalist/utils/token"
)

func (s *Server) credentialsByCustomerId(ctx *gin.Context) {
	var response responser.Response
	var reqParams customerIdRequestParam
	var err error
	err = ctx.ShouldBindUri(&reqParams)
	if err != nil {
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}

	val := ctx.MustGet(middleware.AuthPayloadKey).(*token.Payload)
	if val.CustomerId != *reqParams.CustomerId {
		err = errors.New("you do not have access to this user")
		response = s.responser.New(nil, err, responser.UNAUTHORIZED)
		ctx.JSON(response.Status, response)
		return
	}

	parsedCustomerId, err := uuid.Parse(*reqParams.CustomerId)
	if err != nil {
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}

	credentials, err := s.credentialsDomain.GetCredentials(ctx, parsedCustomerId)
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

	response = s.responser.New(credentials, err, responser.OK)
	ctx.JSON(response.Status, response)
	return
}

func (s *Server) createCredential(ctx *gin.Context) {
	var response responser.Response
	var reqParams customerIdRequestParam
	var request createCredentialRequest
	var err error
	err = ctx.ShouldBindUri(&reqParams)
	if err != nil {
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}

	val := ctx.MustGet(middleware.AuthPayloadKey).(*token.Payload)
	if val.CustomerId != *reqParams.CustomerId {
		err = errors.New("you do not have access to this user")
		response = s.responser.New(nil, err, responser.UNAUTHORIZED)
		ctx.JSON(response.Status, response)
		return
	}

	parsedCustomerId, err := uuid.Parse(*reqParams.CustomerId)
	if err != nil {
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}

	err = s.credentialsDomain.CreateCredential(ctx, credentials.CreateCredentialDto{
		Title:           request.Title,
		LoginName:       request.LoginName,
		Secret:          request.Secret,
		Description:     request.Description,
		ShowImmediately: request.ShowImmediately,
		SendToEmail:     request.SendToEmail,
		SendToPhone:     request.SendToPhone,
		CustomerId:      parsedCustomerId,
	})
	if err != nil {
		response = s.responser.New(nil, err, responser.FAIL)
		ctx.JSON(response.Status, response)
		return
	}

	response = s.responser.New(okResponse{Status: "ok"}, nil, responser.OK)
	ctx.JSON(response.Status, response)
	return
}

func (s *Server) updateCredential(ctx *gin.Context) {
	var response responser.Response
	var reqParams updateCredentialRequestParams
	var request updateCredentialRequest
	var err error
	err = ctx.ShouldBindUri(&reqParams)
	if err != nil {
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}

	val := ctx.MustGet(middleware.AuthPayloadKey).(*token.Payload)
	if val.CustomerId != *reqParams.CustomerId {
		err = errors.New("you do not have access to this user")
		response = s.responser.New(nil, err, responser.UNAUTHORIZED)
		ctx.JSON(response.Status, response)
		return
	}

	parsedCustomerId, err := uuid.Parse(*reqParams.CustomerId)
	if err != nil {
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}
	parsedCredentialId, err := uuid.Parse(*reqParams.CredentialId)
	if err != nil {
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}

	err = s.credentialsDomain.UpdateCredential(ctx, credentials.UpdateCredentialDto{
		CreateCredentialDto: credentials.CreateCredentialDto{
			Title:           request.Title,
			LoginName:       request.LoginName,
			Secret:          request.Secret,
			Description:     request.Description,
			ShowImmediately: request.ShowImmediately,
			SendToEmail:     request.SendToEmail,
			SendToPhone:     request.SendToPhone,
			CustomerId:      parsedCustomerId,
		},
		CredentialId: parsedCredentialId,
	})
	if err != nil {
		response = s.responser.New(nil, err, responser.FAIL)
		ctx.JSON(response.Status, response)
		return
	}

	response = s.responser.New(okResponse{Status: "ok"}, nil, responser.OK)
	ctx.JSON(response.Status, response)
	return
}
