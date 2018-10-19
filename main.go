package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	"WannaChat/controller/userController"
	"WannaChat/controller/groupController"
	"WannaChat/controller/membershipController"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/api")
	{
		// user
		api.POST("/user/signup", userController.Signup)
		api.POST("/user/login", userController.Login)
		api.GET("/user", userController.GetAllUsers)
		api.GET("/user/:id", userController.GetUserById)

		// group
		api.POST("/group/create-group", groupController.CreateGroup)
		api.GET("/group/:id", groupController.GetGroupById)
		api.DELETE("/group/:id", groupController.DeleteGroupById)

		// membership
		api.POST("/membership/create-membership", membershipController.CreateMembership)
	}

	r.Run()
}
