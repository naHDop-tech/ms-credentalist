package api

import (
	"github.com/gin-gonic/gin"
	opt_auth "github.com/naHDop-tech/ms-credentalist/domain/opt-auth"
	"github.com/naHDop-tech/ms-credentalist/utils/responser"
)

func (s *Server) sendOpt(ctx *gin.Context) {
	var response responser.Response
	var request verifyEmailRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		response = s.responser.New(nil, err, responser.BAD_REQUEST)
		ctx.JSON(response.Status, response)
		return
	}

	err = s.optAuthDomain.SentOpt(ctx, opt_auth.CreateOptAuthRecord{
		Email:    request.Email,
		UserName: request.UserName,
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
