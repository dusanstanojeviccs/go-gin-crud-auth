package books

import (
	"go-gin-crud-auth/security"
	"go-gin-crud-auth/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func get(ctx *gin.Context) {
	userId := utils.Session.GetUserId(ctx)

	ctx.JSON(200, LiftRepository.findAll(userId))
}

func getById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	userId := utils.Session.GetUserId(ctx)

	ctx.JSON(200, LiftRepository.findById(id, userId))
}

func post(ctx *gin.Context) {
	userId := utils.Session.GetUserId(ctx)

	lift := new(Lift)

	ctx.BindJSON(lift)

	LiftRepository.create(lift, userId)

	ctx.JSON(200, *lift)
}

func put(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	userId := utils.Session.GetUserId(ctx)

	lift := new(Lift)

	ctx.BindJSON(lift)

	if lift.Id != id {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id missmatch"})
		return
	}

	LiftRepository.update(lift, userId)

	ctx.JSON(200, *lift)
}

func delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	userId := utils.Session.GetUserId(ctx)

	LiftRepository.delete(id, userId)

	ctx.JSON(200, gin.H{})
}
func RegisterEndpoints(server *gin.Engine) {
	server.GET("/lifts", security.LoggedInFilter, get)
	server.GET("/lifts/:id", security.LoggedInFilter, getById)
	server.POST("/lifts", security.LoggedInFilter, post)
	server.PUT("/lifts/:id", security.LoggedInFilter, put)
	server.DELETE("/lifts/:id", security.LoggedInFilter, delete)
}
