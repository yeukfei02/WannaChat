package main

import (
  // "fmt"
  "github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

  "WannaChat/controller/userController"
)

func main() {
  r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/api")
  {
		api.POST("/users/signup", userController.Signup)
		api.POST("/users/login", userController.Login)
		// api.GET("/users", userController.GetAllUsers)
		// api.GET("/users/:id", userController.GetUserById)
		// api.POST("/users/", userController.CreateUser)
		// api.PUT("/users/:id", userController.UpdateUserById)
		// api.DELETE("/users/:id", userController.DeleteUserById)
  }

	r.Run()
}
