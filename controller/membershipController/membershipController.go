package membershipController

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"

	"WannaChat/common"
	"WannaChat/model/membershipModel"
)

// request body
type Membership struct {
	UserId     int64 `json:"userId"`
	GroupId    int64 `json:"groupId"`
	jwt.StandardClaims
}

func CreateMembership(c *gin.Context)  {
  tokenValid := common.CheckAuth(c)
	if tokenValid {
		var membership Membership
		c.BindJSON(&membership)

		membershipModel.InsertMembership(membership.UserId, membership.GroupId)

		c.JSON(200, gin.H{
			"message":         "membership created!",
			"membership":      membership,
			"createdDateTime": time.Now(),
		})
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token!",
		})
	}
}
