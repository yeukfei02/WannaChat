package groupModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Group struct {
	gorm.Model
	GroupId    int    `json:"groupId" gorm:"primary_key; auto_increment;"`
	UserEmail  string `json:"userEmail" gorm:"type:varchar(255); not null;"`
	GroupLabel string `json:"groupLabel" gorm:"type:varchar(255); not null;"`
}
