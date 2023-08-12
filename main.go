package main

import (
	"go-gin-crud-auth/middleware"
	"go-gin-crud-auth/users"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.Use(middleware.Auth)
	users.RegisterEndpoints(server)

	server.Run("localhost:8080")
}
