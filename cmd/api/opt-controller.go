package api

import (
	"github.com/gin-gonic/gin"
	opt_auth "github.com/naHDop-tech/ms-credentalist/domain/opt-auth"
	"github.com/naHDop-tech/ms-credentalist/domain/user"
	"github.com/naHDop-tech/ms-credentalist/utils/responser"
	token2 "github.com/naHDop-tech/ms-credentalist/utils/token"
)

func (s *Server) sendOtp(ctx *gin.Context) {
	var response responser.Response
	var request verifyEmailRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}

	customerId, err := s.userDomain.CreateUser(ctx, user.CreateUserDto{
		Email:    request.Email,
		UserName: request.UserName,
	})
	if err != nil {
		response = s.responser.New(nil, err, responser.FAIL)
		ctx.JSON(response.Status, response)
		return
	}

	err = s.optAuthDomain.SendOpt(ctx, *customerId)
	if err != nil {
		response = s.responser.New(nil, err, responser.FAIL)
		ctx.JSON(response.Status, response)
		return
	}

	response = s.responser.New(okResponse{Status: "ok"}, nil, responser.OK)
	ctx.JSON(response.Status, response)
	return
}

func (s *Server) resendOtp(ctx *gin.Context) {
	var response responser.Response
	var request verifyEmailRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}

	customer, err := s.userDomain.GetCustomerByName(ctx, request.UserName)
	if err != nil {
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}

	err = s.optAuthDomain.SendOpt(ctx, customer.ID)
	if err != nil {
		response = s.responser.New(nil, err, responser.FAIL)
		ctx.JSON(response.Status, response)
		return
	}

	response = s.responser.New(okResponse{Status: "ok"}, nil, responser.OK)
	ctx.JSON(response.Status, response)
	return
}

func (s *Server) verifyOtp(ctx *gin.Context) {
	var response responser.Response
	var request verifyOtp
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}

	err = s.optAuthDomain.VerifyCustomerOtpCode(ctx, opt_auth.VerifyCustomerOtpDto{
		Otp:      request.Otp,
		UserName: request.UserName,
	})
	if err != nil {
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}

	customer, err := s.userDomain.GetCustomerByName(ctx, request.UserName)
	if err != nil {
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}

	token, err := s.tokenMaker.CreateToken(token2.UserPayload{
		CustomerId: customer.ID.String(),
		UserName:   customer.UserName,
	}, s.config.AccessTokenDuration)

	response = s.responser.New(tokenResponse{Token: token}, nil, responser.OK)
	ctx.JSON(response.Status, response)
	return
}
