package membershipModel

import (
	. "WannaChat/common"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Membership struct {
	gorm.Model
	MembershipId int `json:"membershipId" gorm:"primary_key; auto_increment;"`
	UserId       int `json:"userId" gorm:"foreignkey;"`
	GroupId      int `json:"groupId" gorm:"foreignkey;"`
}

func InsertMembership(userId int, groupId int) {
	postgresInfo := GetPostgresInfo()
	db, err := gorm.Open("postgres", postgresInfo)
	CheckErr(err)
	defer db.Close()

	db.AutoMigrate(&Membership{})

	membership := Membership{
		UserId:  userId,
		GroupId: groupId,
	}
	db.Create(&membership)
}
