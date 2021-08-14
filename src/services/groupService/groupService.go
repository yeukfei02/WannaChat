package groupService

import (
	"WannaChat/src/common"
	"WannaChat/src/model/groupModel"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB = common.OpenPostgresDBLazy()

// InsertGroup service
func InsertGroup(groupLabel string) {
	group := groupModel.Group{
		GroupLabel: groupLabel,
	}
	db.Create(&group)
}

// GetAllGroups service
func GetAllGroups() []groupModel.Group {
	var groups []groupModel.Group
	db.Find(&groups)
	return groups
}

// GetGroupByID service
func GetGroupByID(groupID string) groupModel.Group {
	var group groupModel.Group
	db.Where("ID = ?", groupID).Find(&group)
	return group
}

// DeleteGroupByID service
func DeleteGroupByID(groupID string) {
	db.Where("ID = ?", groupID).Delete(&groupModel.Group{})
}
