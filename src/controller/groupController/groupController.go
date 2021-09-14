package groupController

import (
	"github.com/gin-gonic/gin"

	"WannaChat/src/helpers"
	"WannaChat/src/services/groupService"
)

// Group request body
type Group struct {
	GroupLabel string `json:"groupLabel"`
}

// CreateGroup controller
func CreateGroup(c *gin.Context) {
	tokenValid := helpers.CheckAuth(c)
	if tokenValid {
		var group Group
		c.BindJSON(&group)

		groupService.InsertGroup(group.GroupLabel)

		c.JSON(201, gin.H{
			"message": "group created",
		})
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token",
		})
	}
}

// GetAllGroups controller
func GetAllGroups(c *gin.Context) {
	tokenValid := helpers.CheckAuth(c)
	if tokenValid {
		groupsList := groupService.GetAllGroups()

		c.JSON(200, gin.H{
			"message": "get all groups",
			"groups":  groupsList,
		})
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token",
		})
	}
}

// GetGroupByID controller
func GetGroupByID(c *gin.Context) {
	tokenValid := helpers.CheckAuth(c)
	if tokenValid {
		groupID := c.Param("id")
		group := groupService.GetGroupByID(groupID)

		c.JSON(200, gin.H{
			"message": "get group by id",
			"group":   group,
		})
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token",
		})
	}
}

// DeleteGroupByID controller
func DeleteGroupByID(c *gin.Context) {
	tokenValid := helpers.CheckAuth(c)
	if tokenValid {
		groupID := c.Param("id")
		groupService.DeleteGroupByID(groupID)

		c.JSON(200, gin.H{
			"message": "delete group by id",
		})
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token",
		})
	}
}
