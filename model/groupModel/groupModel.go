package groupModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	. "WannaChat/common"
)

type Group struct {
	gorm.Model
	GroupLabel string `json:"groupLabel" gorm:"type:varchar(255); not null;"`
}

func InsertGroup(groupLabel string) {
	db := OpenPostgresDBLazy()
	defer db.Close()
	CheckTableExists(db, &Group{})

	group := Group{
		GroupLabel: groupLabel,
	}
	db.Create(&group)
}

func GetAllGroups() []Group {
	db := OpenPostgresDBLazy()
	defer db.Close()
	CheckTableExists(db, &Group{})

	var groups []Group
	db.Find(&groups)
	return groups
}

func GetGroupById(groupId string) Group {
	db := OpenPostgresDBLazy()
	defer db.Close()
	CheckTableExists(db, &Group{})

	var group Group
	db.Where("ID = ?", groupId).Find(&group)
	return group
}

func DeleteGroupById(groupId string) {
	db := OpenPostgresDBLazy()
	defer db.Close()
	CheckTableExists(db, &Group{})

	db.Where("ID = ?", groupId).Delete(&Group{})
}
