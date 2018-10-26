package membershipModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"WannaChat/common"
)

type Membership struct {
	gorm.Model
	UserId  uint `json:"userId" gorm:"foreignkey;"`
	GroupId uint `json:"groupId" gorm:"foreignkey;"`
}

func InsertMembership(userId uint, groupId uint) {
	db := common.OpenPostgresDBLazy()
	defer db.Close()
	db.AutoMigrate(&Membership{})

	membership := Membership{
		UserId:  userId,
		GroupId: groupId,
	}
	db.Create(&membership)
}

func GetAllMemberships() []Membership {
	db := common.OpenPostgresDBLazy()
	defer db.Close()
	db.AutoMigrate(&Membership{})

	var memberships []Membership
	db.Find(&memberships)
	return memberships
}
