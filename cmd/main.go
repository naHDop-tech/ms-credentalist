package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/naHDop-tech/ms-credentalist/cmd/api"
	opt_auth "github.com/naHDop-tech/ms-credentalist/domain/opt-auth"
	"github.com/naHDop-tech/ms-credentalist/domain/user"
	"github.com/naHDop-tech/ms-credentalist/utils"
)

func main() {
	conf, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not read from config:", err)
	}
	dbPort, err := strconv.Atoi(conf.DBPort)
	if err != nil {
		log.Fatal("Could not parse server port:", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.DBHost, dbPort, conf.DBUser, conf.DBPassword, conf.DBName)

	conn, err := sql.Open(conf.DBDriver, psqlInfo)
	if err != nil {
		log.Fatal("Could not connect to db:", err)
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	serverAddress := fmt.Sprintf("%s:%s", conf.AppHost, conf.AppPort)
	optDomain := opt_auth.NewOptAuthDomain(conn)
	userDomain := user.NewUserDomain(conn)

	server, err := api.NewServer(conn, conf, optDomain, userDomain)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Server has not started", err)
	}
}
