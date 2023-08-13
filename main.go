package main

import (
	"go-gin-crud-auth/app/lifts"
	"go-gin-crud-auth/app/users"
	"go-gin-crud-auth/middleware"
	"go-gin-crud-auth/utils"
	"go-gin-crud-auth/utils/db"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.Config.Init()

	db.Init()

	server := gin.Default()

	server.Use(middleware.Auth)
	users.RegisterEndpoints(server)
	lifts.RegisterEndpoints(server)

	server.Run(utils.Config.Server.Address)
}
