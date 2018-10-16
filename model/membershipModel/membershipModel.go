package membershipModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Membership struct {
	UserEmail string `json:"userEmail"`
	GroupID   string `json:"groupId"`
}
