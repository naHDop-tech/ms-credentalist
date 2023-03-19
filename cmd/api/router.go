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

	v1.POST("/send-opt", s.sendOpt)
	v1.POST("/re-send-opt", s.resendOpt)
	v1.POST("/user")
	v1.POST("/login")

	v1AuthGroupRoute := v1.Group("/").Use(middleware.AuthMiddleware(s.tokenMaker))
	{
		v1AuthGroupRoute.GET("/user/:user_id")
	}

	s.router = router
}
