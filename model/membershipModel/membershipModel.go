package membershipModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Membership struct {
	gorm.Model
	MembershipId 	int     `json:"membershipId" gorm:"primary_key; auto_increment;"`
	UserEmail 		string 	`json:"userEmail" gorm:"type:varchar(255); not null;"`
	UserId   			int 		`json:"userId" gorm:"foreignkey;"`
	GroupId   		int 		`json:"groupId" gorm:"foreignkey;"`
}
