package groupModel

import (
	"WannaChat/common"
	"WannaChat/schema"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB = common.OpenPostgresDBLazy()

// InsertGroup model
func InsertGroup(groupLabel string) {
	group := schema.Group{
		GroupLabel: groupLabel,
	}
	db.Create(&group)
}

// GetAllGroups model
func GetAllGroups() []schema.Group {
	var groups []schema.Group
	db.Find(&groups)
	return groups
}

// GetGroupByID model
func GetGroupByID(groupID string) schema.Group {
	var group schema.Group
	db.Where("ID = ?", groupID).Find(&group)
	return group
}

// DeleteGroupByID model
func DeleteGroupByID(groupID string) {
	db.Where("ID = ?", groupID).Delete(&schema.Group{})
}
