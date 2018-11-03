package membershipController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"

	. "WannaChat/common"
	"WannaChat/model/groupModel"
	"WannaChat/model/membershipModel"
	"WannaChat/model/userModel"
)

// request body
type Membership struct {
	UserId  uint `json:"userId"`
	GroupId uint `json:"groupId"`
}

func CreateMembership(c *gin.Context) {
	tokenValid := CheckAuth(c)
	if tokenValid {
		var membership Membership
		c.BindJSON(&membership)

		user := userModel.GetUserById(fmt.Sprint(membership.UserId))
		group := groupModel.GetGroupById(fmt.Sprint(membership.GroupId))
		fmt.Println(user.ID, group.ID)

		if user.ID > 0 && group.ID > 0 {
			membershipModel.InsertMembership(membership.UserId, membership.GroupId)

			c.JSON(200, gin.H{
				"message":         "membership created!",
				"membership":      membership,
				"createdDateTime": time.Now(),
			})
		} else {
			c.JSON(400, gin.H{
				"message": "userId or groupId not found!",
			})
		}
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token!",
		})
	}
}

func GetAllMemberships(c *gin.Context) {
	tokenValid := CheckAuth(c)
	if tokenValid {
		membershipsList := membershipModel.GetAllMemberships()

		c.JSON(200, gin.H{
			"message":     "get all memberships!",
			"memberships": membershipsList,
			"count":       len(membershipsList),
		})
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token!",
		})
	}
}
