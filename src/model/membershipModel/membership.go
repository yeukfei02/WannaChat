package membershipModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Membership model
type Membership struct {
	gorm.Model
	UserFk  uint `json:"userFk" gorm:"foreignkey;"`
	GroupFk uint `json:"groupFk" gorm:"foreignkey;"`
}
