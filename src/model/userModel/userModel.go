package userModel

import (
	"WannaChat/src/common"
	"WannaChat/src/schema"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB = common.OpenPostgresDBLazy()

// InsertUser model
func InsertUser(email string, password string) {
	user := schema.User{
		Email:    email,
		Password: password,
	}
	db.Create(&user)
}

// GetUserPassword model
func GetUserPassword(email string) string {
	var user schema.User
	db.Where("email = ?", email).Find(&user)
	return user.Password
}

// GetAllUsers model
func GetAllUsers() []schema.User {
	var users []schema.User
	db.Find(&users)
	return users
}

// GetUserByID model
func GetUserByID(userID string) schema.User {
	var user schema.User
	db.Where("ID = ?", userID).Find(&user)
	return user
}
