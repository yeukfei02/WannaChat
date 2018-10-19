package membershipModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"WannaChat/common"
)

type Membership struct {
	gorm.Model
	UserId  int64 `json:"userId" gorm:"foreignkey;"`
	GroupId int64 `json:"groupId" gorm:"foreignkey;"`
}

func InsertMembership(userId int64, groupId int64) {
	postgresInfo := common.GetPostgresInfo()
	db, err := gorm.Open("postgres", postgresInfo)
	common.CheckErr(err)
	defer db.Close()
	db.AutoMigrate(&Membership{})

	membership := Membership{
		UserId:  userId,
		GroupId: groupId,
	}
	db.Create(&membership)
}
