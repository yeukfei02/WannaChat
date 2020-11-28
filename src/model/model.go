package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User model
type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"type:varchar(255); unique; not null;"`
	Password string `json:"password" gorm:"type:varchar(255); not null;"`
}

// Group model
type Group struct {
	gorm.Model
	GroupLabel string `json:"groupLabel" gorm:"type:varchar(255); not null;"`
}

// Membership model
type Membership struct {
	gorm.Model
	UserFk  uint `json:"userFk" gorm:"foreignkey;"`
	GroupFk uint `json:"groupFk" gorm:"foreignkey;"`
}
