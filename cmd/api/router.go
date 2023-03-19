package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/naHDop-tech/ms-credentalist/cmd/middleware"
)

func (s *Server) SetupRouter() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "PATCH", "GET"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Referer", "User-Agent"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	v1 := router.Group("/api/v1")

	v1.POST("/user")
	v1.POST("/login")

	v1AuthGroupRoute := v1.Group("/").Use(middleware.AuthMiddleware(s.tokenMaker))
	{
		v1AuthGroupRoute.GET("/user/:user_id")
	}

	s.router = router
}
