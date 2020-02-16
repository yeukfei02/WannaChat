package membershipRoutes

import (
	"github.com/gin-gonic/gin"

	"WannaChat/src/controller/membershipController"
)

// Routes Membership
func Routes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/membership/create-membership", membershipController.CreateMembership)
		api.GET("/membership", membershipController.GetAllMemberships)
		api.GET("/membership/get-membership-by-group-id", membershipController.GetMembershipsByGroupID)
	}
}
