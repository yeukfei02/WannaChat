package userModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"WannaChat/common"
)

// User model
type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"type:varchar(255); unique; not null;"`
	Password string `json:"password" gorm:"type:varchar(255); not null;"`
}

// InsertUser model
func InsertUser(email string, password string) {
	db := common.OpenPostgresDBLazy()
	defer db.Close()
	common.CheckTableExists(db, &User{})

	user := User{
		Email:    email,
		Password: password,
	}
	db.Create(&user)
}

// GetUserPassword model
func GetUserPassword(email string) string {
	db := common.OpenPostgresDBLazy()
	defer db.Close()
	common.CheckTableExists(db, &User{})

	var user User
	db.Where("email = ?", email).Find(&user)
	return user.Password
}

// GetAllUsers model
func GetAllUsers() []User {
	db := common.OpenPostgresDBLazy()
	defer db.Close()
	common.CheckTableExists(db, &User{})

	var users []User
	db.Find(&users)
	return users
}

// GetUserByID model
func GetUserByID(userID string) User {
	db := common.OpenPostgresDBLazy()
	defer db.Close()
	common.CheckTableExists(db, &User{})

	var user User
	db.Where("ID = ?", userID).Find(&user)
	return user
}
