package main

import (
	"Assignment4/handler"
	"Assignment4/middleware"
	"Assignment4/model"

	"github.com/gin-gonic/gin"
)

func main() {
	model.SeedUsers()
	router := gin.Default()

	auth := router.Group("/api/auth")
	{
		auth.POST("/login", handler.Login)
		auth.GET("/me", middleware.Authenticate, handler.Me)
	}


	router.Run(":8080")

}
