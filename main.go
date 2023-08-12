package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/user", func(ctx *gin.Context) {
		user := User{Name: "John"}
		ctx.JSON(200, &user)
	})

	server.Run("localhost:8080")
}
