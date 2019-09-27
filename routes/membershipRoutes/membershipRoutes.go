package membershipRoutes

import (
	"github.com/gin-gonic/gin"

	"WannaChat/controller/membershipController"
)

func Routes(r *gin.Engine) {
	api := r.Group("/api")
	{
		// membership
		api.POST("/membership/create-membership", membershipController.CreateMembership)
		api.GET("/membership", membershipController.GetAllMemberships)
		api.GET("/membership/get-membership-by-group-id", membershipController.GetMembershipsByGroupId)
	}
}
