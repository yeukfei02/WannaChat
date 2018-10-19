package groupController

import (
	"WannaChat/common"
	"WannaChat/model/groupModel"
	"WannaChat/model/membershipModel"
	"WannaChat/model/userModel"
	. "github.com/dgrijalva/jwt-go"
	. "github.com/gin-gonic/gin"
)

// request body
type Group struct {
	Email      string `json:"email"`
	GroupLabel string `json:"group_label"`
	StandardClaims
}

//Auth Required
func CreateNewGroup(c *Context) {
	tokenValid := common.CheckAuth(c)
	if tokenValid {
		var group Group
		c.BindJSON(&group)

		var user = userModel.GetUserByEmail(group.Email)
		var groupID = groupModel.InsertGroup(group.GroupLabel)
		membershipModel.InsertMembership(user.UserId, groupID)

		c.JSON(200, H{
			"message": "group created!",
		})
	} else {
		c.JSON(404, H{
			"message": "wrong or missing token!",
		})
	}
}
