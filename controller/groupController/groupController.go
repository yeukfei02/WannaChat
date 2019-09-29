package groupController

import (
	"time"

	"github.com/gin-gonic/gin"

	"WannaChat/common"
	"WannaChat/model/groupModel"
)

// Group request body
type Group struct {
	GroupLabel string `json:"groupLabel"`
}

// CreateGroup controller
func CreateGroup(c *gin.Context) {
	tokenValid := common.CheckAuth(c)
	if tokenValid {
		var group Group
		c.BindJSON(&group)

		groupModel.InsertGroup(group.GroupLabel)

		c.JSON(201, gin.H{
			"message":         "group created!",
			"group":           group,
			"createdDateTime": time.Now(),
		})
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token!",
		})
	}
}

// GetAllGroups controller
func GetAllGroups(c *gin.Context) {
	tokenValid := common.CheckAuth(c)
	if tokenValid {
		groupsList := groupModel.GetAllGroups()

		c.JSON(200, gin.H{
			"message": "get all groups!",
			"groups":  groupsList,
			"count":   len(groupsList),
		})
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token!",
		})
	}
}

// GetGroupByID controller
func GetGroupByID(c *gin.Context) {
	tokenValid := common.CheckAuth(c)
	if tokenValid {
		groupID := c.Param("id")
		group := groupModel.GetGroupByID(groupID)

		c.JSON(200, gin.H{
			"message": "get group by id!",
			"group":   group,
		})
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token!",
		})
	}
}

// DeleteGroupByID controller
func DeleteGroupByID(c *gin.Context) {
	tokenValid := common.CheckAuth(c)
	if tokenValid {
		groupID := c.Param("id")
		groupModel.DeleteGroupByID(groupID)

		c.JSON(200, gin.H{
			"message": "delete group by id!",
		})
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token!",
		})
	}
}
