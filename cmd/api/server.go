package api

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/naHDop-tech/ms-credentalist/utils"
	"github.com/naHDop-tech/ms-credentalist/utils/responser"
	"github.com/naHDop-tech/ms-credentalist/utils/token"
)

type Server struct {
	router     *gin.Engine
	tokenMaker token.Maker
	connect    *sql.DB
	config     utils.Config
	responser  responser.Responser
}

func NewServer(
	connect *sql.DB,
	config utils.Config,
) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %s", err)
	}
	server := &Server{
		tokenMaker: tokenMaker,
		connect:    connect,
		config:     config,
		responser:  responser.NewResponser(),
	}

	server.SetupRouter()
	return server, nil
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
