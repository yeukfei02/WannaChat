package membershipModel

import (
	"WannaChat/src/common"
	"WannaChat/src/schema"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB = common.OpenPostgresDBLazy()

// InsertMembership model
func InsertMembership(userFk uint, groupFk uint) {
	membership := schema.Membership{
		UserFk:  userFk,
		GroupFk: groupFk,
	}
	db.Create(&membership)
}

// GetAllMemberships model
func GetAllMemberships() []schema.Membership {
	var memberships []schema.Membership
	db.Find(&memberships)
	return memberships
}

// GetMembershipsByGroupID model
func GetMembershipsByGroupID(groupID string) []schema.Membership {
	var memberships []schema.Membership
	db.Where("group_fk = ?", groupID).Find(&memberships)
	return memberships
}
