package groupRoutes

import (
	"github.com/gin-gonic/gin"

	"WannaChat/src/controller/groupController"
)

// Routes Group
func Routes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/group/create-group", groupController.CreateGroup)
		api.GET("/group", groupController.GetAllGroups)
		api.GET("/group/:id", groupController.GetGroupByID)
		api.DELETE("/group/:id", groupController.DeleteGroupByID)
	}
}
