package mainRoutes

import (
	"github.com/gin-gonic/gin"

	"WannaChat/src/controller/mainController"
)

// Routes Group
func Routes(r *gin.Engine) {
	api := r.Group("/")
	{
		api.GET("", mainController.GetMain)
	}
}
