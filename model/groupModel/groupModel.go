package groupModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"WannaChat/common"
)

type Group struct {
	gorm.Model
	GroupLabel string `json:"groupLabel" gorm:"type:varchar(255); not null;"`
}

func InsertGroup(groupLabel string) {
	db := common.OpenPostgresDBLazy()
	defer db.Close()
	db.AutoMigrate(&Group{})

	group := Group{
		GroupLabel: groupLabel,
	}
	db.Create(&group)
}

func GetGroupById(groupId string) Group {
	db := common.OpenPostgresDBLazy()
	defer db.Close()
	db.AutoMigrate(&Group{})

	var group Group
	db.Where("ID = ?", groupId).Find(&group)
	return group
}

func DeleteGroupById(groupId string) {
	db := common.OpenPostgresDBLazy()
	defer db.Close()
	db.AutoMigrate(&Group{})

	db.Where("ID = ?", groupId).Delete(&Group{})
}
