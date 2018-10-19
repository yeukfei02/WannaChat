package groupModel

import (
	"WannaChat/common"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Group struct {
	gorm.Model
	GroupId    int    `json:"groupId" gorm:"primary_key; auto_increment;"`
	GroupLabel string `json:"groupLabel" gorm:"type:varchar(255); not null;"`
}

func InsertGroup(label string) int {
	postgresInfo := common.GetPostgresInfo()
	db, err := gorm.Open("postgres", postgresInfo)
	checkErr(err)
	defer db.Close()

	db.AutoMigrate(&Group{})

	group := Group{
		GroupLabel: label,
	}
	db.Create(&group)
	return group.GroupId
}

func DeleteGroup(groupId string) {
	postgresInfo := common.GetPostgresInfo()
	db, err := gorm.Open("postgres", postgresInfo)
	checkErr(err)
	defer db.Close()

	db.AutoMigrate(&Group{})

	db.Where("groupId = ?", groupId).Delete(&Group{})
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("error = " + err.Error())
	}
}
