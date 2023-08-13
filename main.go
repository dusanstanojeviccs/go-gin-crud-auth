package main

import (
	"go-gin-crud-auth/app/lifts"
	"go-gin-crud-auth/app/users"
	"go-gin-crud-auth/middleware"
	"go-gin-crud-auth/utils/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()

	server := gin.Default()

	server.Use(middleware.Auth)
	users.RegisterEndpoints(server)
	lifts.RegisterEndpoints(server)

	server.Run("localhost:8080")
}
