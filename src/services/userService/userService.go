package userService

import (
	"WannaChat/src/common"
	"WannaChat/src/model/userModel"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB = common.OpenPostgresDBLazy()

// InsertUser service
func InsertUser(email string, password string) {
	user := userModel.User{
		Email:    email,
		Password: password,
	}
	db.Create(&user)
}

// GetUserPassword service
func GetUserPassword(email string) string {
	var user userModel.User
	db.Where("email = ?", email).Find(&user)
	return user.Password
}

// GetAllUsers service
func GetAllUsers() []userModel.User {
	var users []userModel.User
	db.Find(&users)
	return users
}

// GetUserByID service
func GetUserByID(userID string) userModel.User {
	var user userModel.User
	db.Where("ID = ?", userID).Find(&user)
	return user
}
