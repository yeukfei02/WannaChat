package membershipService

import (
	"WannaChat/src/helpers"
	"WannaChat/src/model/membershipModel"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB = helpers.OpenPostgresDBLazy()

// InsertMembership service
func InsertMembership(userFk uint, groupFk uint) {
	membership := membershipModel.Membership{
		UserFk:  userFk,
		GroupFk: groupFk,
	}
	db.Create(&membership)
}

// GetAllMemberships service
func GetAllMemberships() []membershipModel.Membership {
	var memberships []membershipModel.Membership
	db.Find(&memberships)
	return memberships
}

// GetMembershipsByGroupID service
func GetMembershipsByGroupID(groupID string) []membershipModel.Membership {
	var memberships []membershipModel.Membership
	db.Where("group_fk = ?", groupID).Find(&memberships)
	return memberships
}
