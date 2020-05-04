package mainController

import (
	"github.com/gin-gonic/gin"
)

// GetMain controller
func GetMain(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "WannaChat api",
	})
}
