package membershipModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Membership struct {
	gorm.Model
	UserEmail string `json:"user_email"`
	GroupID   string `json:"group_id"`
}
