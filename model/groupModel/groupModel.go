package groupModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Group struct {
	gorm.Model
	UserEmail  string `json:"user_email"`
	GroupLabel string `json:"group_label"`
}
