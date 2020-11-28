package membershipService

import (
	"WannaChat/src/common"
	"WannaChat/src/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB = common.OpenPostgresDBLazy()

// InsertMembership service
func InsertMembership(userFk uint, groupFk uint) {
	membership := model.Membership{
		UserFk:  userFk,
		GroupFk: groupFk,
	}
	db.Create(&membership)
}

// GetAllMemberships service
func GetAllMemberships() []model.Membership {
	var memberships []model.Membership
	db.Find(&memberships)
	return memberships
}

// GetMembershipsByGroupID service
func GetMembershipsByGroupID(groupID string) []model.Membership {
	var memberships []model.Membership
	db.Where("group_fk = ?", groupID).Find(&memberships)
	return memberships
}
