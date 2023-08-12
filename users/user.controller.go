package users

import (
	"go-gin-crud-auth/security"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func get(ctx *gin.Context) {
	ctx.JSON(200, UserRepository.findAll())
}

func getById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	ctx.JSON(200, UserRepository.findById(id))
}

func post(ctx *gin.Context) {
	user := new(User)

	ctx.BindJSON(user)

	UserRepository.create(user)

	ctx.JSON(200, *user)
}

func put(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	user := new(User)

	ctx.BindJSON(user)

	if user.Id != id {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id missmatch"})
		return
	}

	UserRepository.update(user)

	ctx.JSON(200, *user)
}

func delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	UserRepository.delete(id)

	ctx.JSON(200, gin.H{})
}
func RegisterEndpoints(server *gin.Engine) {
	server.GET("/users", get)
	server.GET("/users/:id", getById)
	server.POST("/users", post)
	server.PUT("/users/:id", security.LoggedInFilter, put)
	server.DELETE("/users/:id", security.LoggedInFilter, delete)
}
