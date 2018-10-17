package membershipModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Membership struct {
	gorm.Model
	MembershipId 	int     `gorm:"primary_key; auto_increment;"`
	UserEmail 		string 	`gorm:"type:varchar(255); not null;"`
	UserId   			int 		`gorm:"foreignkey;"`
	GroupId   		int 		`gorm:"foreignkey;"`
}
