package membershipController

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	. "WannaChat/common"
	"WannaChat/model/groupModel"
	"WannaChat/model/membershipModel"
	"WannaChat/model/userModel"
)

// request body
type Membership struct {
	UserFk  uint `json:"userFk"`
	GroupFk uint `json:"groupFk"`
}

func CreateMembership(c *gin.Context) {
	tokenValid := CheckAuth(c)
	if tokenValid {
		var membership Membership
		c.BindJSON(&membership)

		user := userModel.GetUserById(fmt.Sprint(membership.UserFk))
		group := groupModel.GetGroupById(fmt.Sprint(membership.GroupFk))
		fmt.Println("userId = ", user.ID)
		fmt.Println("groupId = ", group.ID)

		if user.ID > 0 && group.ID > 0 {
			membershipModel.InsertMembership(membership.UserFk, membership.GroupFk)

			c.JSON(201, gin.H{
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

func GetMembershipsByGroupId(c *gin.Context) {
	tokenValid := CheckAuth(c)
	if tokenValid {
		groupID := c.Query("groupId")
		membershipsList := membershipModel.GetMembershipsByGroupId(groupID)

		c.JSON(200, gin.H{
			"message":     "get memberships by group id!",
			"memberships": membershipsList,
			"count":       len(membershipsList),
		})
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token!",
		})
	}
}
