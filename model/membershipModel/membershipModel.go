package membershipModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	. "WannaChat/common"
)

type Membership struct {
	gorm.Model
	UserFk  uint `json:"userFk" gorm:"foreignkey;"`
	GroupFk uint `json:"groupFk" gorm:"foreignkey;"`
}

func InsertMembership(userFk uint, groupFk uint) {
	db := OpenPostgresDBLazy()
	defer db.Close()
	CheckTableExists(db, &Membership{})

	membership := Membership{
		UserFk:  userFk,
		GroupFk: groupFk,
	}
	db.Create(&membership)
}

func GetAllMemberships() []Membership {
	db := OpenPostgresDBLazy()
	defer db.Close()
	CheckTableExists(db, &Membership{})

	var memberships []Membership
	db.Find(&memberships)
	return memberships
}

func GetMembershipsByGroupId(groupID string) []Membership {
	db := OpenPostgresDBLazy()
	defer db.Close()
	CheckTableExists(db, &Membership{})

	var memberships []Membership
	db.Where("group_fk = ?", groupID).Find(&memberships)
	return memberships
}
