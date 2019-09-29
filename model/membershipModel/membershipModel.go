package membershipModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"WannaChat/common"
)

// Membership model
type Membership struct {
	gorm.Model
	UserFk  uint `json:"userFk" gorm:"foreignkey;"`
	GroupFk uint `json:"groupFk" gorm:"foreignkey;"`
}

// InsertMembership model
func InsertMembership(userFk uint, groupFk uint) {
	db := common.OpenPostgresDBLazy()
	defer db.Close()
	common.CheckTableExists(db, &Membership{})

	membership := Membership{
		UserFk:  userFk,
		GroupFk: groupFk,
	}
	db.Create(&membership)
}

// GetAllMemberships model
func GetAllMemberships() []Membership {
	db := common.OpenPostgresDBLazy()
	defer db.Close()
	common.CheckTableExists(db, &Membership{})

	var memberships []Membership
	db.Find(&memberships)
	return memberships
}

// GetMembershipsByGroupID model
func GetMembershipsByGroupID(groupID string) []Membership {
	db := common.OpenPostgresDBLazy()
	defer db.Close()
	common.CheckTableExists(db, &Membership{})

	var memberships []Membership
	db.Where("group_fk = ?", groupID).Find(&memberships)
	return memberships
}
