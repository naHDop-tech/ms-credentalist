package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/naHDop-tech/ms-credentalist/cmd/middleware"
)

func (s *Server) SetupRouter() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     s.config.AllowedOrigin,
		AllowMethods:     s.config.AllowedMethods,
		AllowHeaders:     s.config.AllowedHeaders,
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	v1 := router.Group("/api/v1")

	v1.POST("/otp/send", s.sendOtp)
	v1.POST("/otp/re-send", s.resendOtp)
	v1.POST("/otp/verify", s.verifyOtp)

	v1A := v1.Group("/").Use(middleware.AuthMiddleware(s.tokenMaker))
	{
		v1A.GET("/customer/:customer_id", s.customerById)
		v1A.GET("/customer/:customer_id/credentials", s.credentialsByCustomerId)

		v1A.POST("/customer/:customer_id/credential")

		v1A.PATCH("/customer/:customer_id/credential/:credential_id")
	}

	s.router = router
}
