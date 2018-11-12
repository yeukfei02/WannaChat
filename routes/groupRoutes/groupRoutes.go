package groupRoutes

import (
	"github.com/gin-gonic/gin"

	"WannaChat/controller/groupController"
)

func Routes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// group
		api.POST("/group/create-group", groupController.CreateGroup)
		api.GET("/group", groupController.GetAllGroups)
		api.GET("/group/:id", groupController.GetGroupById)
		api.DELETE("/group/:id", groupController.DeleteGroupById)
	}
}
