package main

import (
	"go-gin-crud-auth/users"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	users.RegisterEndpoints(server)

	server.Run("localhost:8080")
}
