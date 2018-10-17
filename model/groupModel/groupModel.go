package groupModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Group struct {
	gorm.Model
	GroupId    int    `gorm:"primary_key; auto_increment;"`
	UserEmail  string `gorm:"type:varchar(255); not null;"`
	GroupLabel string `gorm:"type:varchar(255); not null;"`
}
