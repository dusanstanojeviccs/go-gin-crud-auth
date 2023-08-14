package main

import (
	"go-gin-crud-auth/app/lifts"
	"go-gin-crud-auth/app/users"
	"go-gin-crud-auth/middleware"
	"go-gin-crud-auth/utils"
	"go-gin-crud-auth/utils/db"
	"go-gin-crud-auth/utils/jwt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.Config.Init()

	db.Init()
	jwt.Init()

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:8080" || origin == "http://localhost:5173"
		},
		MaxAge: 12 * time.Hour,
	}))

	server.Use(middleware.Transactional)
	server.Use(middleware.Auth)

	users.RegisterEndpoints(server)
	lifts.RegisterEndpoints(server)

	server.Run(utils.Config.Server.Address)
}
