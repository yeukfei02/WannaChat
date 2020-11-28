package groupService

import (
	"WannaChat/src/common"
	"WannaChat/src/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB = common.OpenPostgresDBLazy()

// InsertGroup service
func InsertGroup(groupLabel string) {
	group := model.Group{
		GroupLabel: groupLabel,
	}
	db.Create(&group)
}

// GetAllGroups service
func GetAllGroups() []model.Group {
	var groups []model.Group
	db.Find(&groups)
	return groups
}

// GetGroupByID service
func GetGroupByID(groupID string) model.Group {
	var group model.Group
	db.Where("ID = ?", groupID).Find(&group)
	return group
}

// DeleteGroupByID service
func DeleteGroupByID(groupID string) {
	db.Where("ID = ?", groupID).Delete(&model.Group{})
}
