package userRoutes

import (
	"github.com/gin-gonic/gin"

	"WannaChat/controller/userController"
)

// Routes User
func Routes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/user/signup", userController.Signup)
		api.POST("/user/login", userController.Login)
		api.GET("/user", userController.GetAllUsers)
		api.GET("/user/:id", userController.GetUserByID)
	}
}
