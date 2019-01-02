package groupController

import (
	"time"

	"github.com/gin-gonic/gin"

	. "WannaChat/common"
	"WannaChat/model/groupModel"
)

// request body
type Group struct {
	GroupLabel string `json:"groupLabel"`
}

func CreateGroup(c *gin.Context) {
	tokenValid := CheckAuth(c)
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

func GetAllGroups(c *gin.Context) {
	tokenValid := CheckAuth(c)
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

func GetGroupById(c *gin.Context) {
	tokenValid := CheckAuth(c)
	if tokenValid {
		groupId := c.Param("id")
		group := groupModel.GetGroupById(groupId)

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

func DeleteGroupById(c *gin.Context) {
	tokenValid := CheckAuth(c)
	if tokenValid {
		groupId := c.Param("id")
		groupModel.DeleteGroupById(groupId)

		c.JSON(200, gin.H{
			"message": "delete group by id!",
		})
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token!",
		})
	}
}
