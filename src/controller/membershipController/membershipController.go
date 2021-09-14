package membershipController

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"WannaChat/src/helpers"
	"WannaChat/src/services/groupService"
	"WannaChat/src/services/membershipService"
	"WannaChat/src/services/userService"
)

// Membership request body
type Membership struct {
	UserFk  uint `json:"userFk"`
	GroupFk uint `json:"groupFk"`
}

// CreateMembership controller
func CreateMembership(c *gin.Context) {
	tokenValid := helpers.CheckAuth(c)
	if tokenValid {
		var membership Membership
		c.BindJSON(&membership)

		user := userService.GetUserByID(fmt.Sprint(membership.UserFk))
		group := groupService.GetGroupByID(fmt.Sprint(membership.GroupFk))
		fmt.Println("userId = ", user.ID)
		fmt.Println("groupId = ", group.ID)

		if user.ID > 0 && group.ID > 0 {
			membershipService.InsertMembership(membership.UserFk, membership.GroupFk)

			c.JSON(201, gin.H{
				"message": "membership created",
			})
		} else {
			c.JSON(400, gin.H{
				"message": "userId or groupId not found!",
			})
		}
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token",
		})
	}
}

// GetAllMemberships controller
func GetAllMemberships(c *gin.Context) {
	tokenValid := helpers.CheckAuth(c)
	if tokenValid {
		membershipsList := membershipService.GetAllMemberships()

		c.JSON(200, gin.H{
			"message":     "get all memberships",
			"memberships": membershipsList,
		})
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token",
		})
	}
}

// GetMembershipsByGroupID controller
func GetMembershipsByGroupID(c *gin.Context) {
	tokenValid := helpers.CheckAuth(c)
	if tokenValid {
		groupID := c.Query("groupId")
		membershipsList := membershipService.GetMembershipsByGroupID(groupID)

		c.JSON(200, gin.H{
			"message":     "get memberships by group id",
			"memberships": membershipsList,
		})
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token",
		})
	}
}
