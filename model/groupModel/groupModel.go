package groupModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Group struct {
	UserEmail  string `json:"userEmail"`
	GroupLabel string `json:"groupLabel"`
}
