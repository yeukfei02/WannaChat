package groupModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"WannaChat/common"
)

// Group model
type Group struct {
	gorm.Model
	GroupLabel string `json:"groupLabel" gorm:"type:varchar(255); not null;"`
}

// InsertGroup model
func InsertGroup(groupLabel string) {
	db := common.OpenPostgresDBLazy()
	defer db.Close()
	common.CheckTableExists(db, &Group{})

	group := Group{
		GroupLabel: groupLabel,
	}
	db.Create(&group)
}

// GetAllGroups model
func GetAllGroups() []Group {
	db := common.OpenPostgresDBLazy()
	defer db.Close()
	common.CheckTableExists(db, &Group{})

	var groups []Group
	db.Find(&groups)
	return groups
}

// GetGroupByID model
func GetGroupByID(groupID string) Group {
	db := common.OpenPostgresDBLazy()
	defer db.Close()
	common.CheckTableExists(db, &Group{})

	var group Group
	db.Where("ID = ?", groupID).Find(&group)
	return group
}

// DeleteGroupByID model
func DeleteGroupByID(groupID string) {
	db := common.OpenPostgresDBLazy()
	defer db.Close()
	common.CheckTableExists(db, &Group{})

	db.Where("ID = ?", groupID).Delete(&Group{})
}
