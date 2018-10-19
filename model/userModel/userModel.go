package userModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"WannaChat/common"
)

type Users []User

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"type:varchar(255); not null;"`
	Password string `json:"password" gorm:"type:varchar(255); not null;"`
}

func InsertUser(email string, password string) {
	postgresInfo := common.GetPostgresInfo()
	db, err := gorm.Open("postgres", postgresInfo)
	common.CheckErr(err)
	defer db.Close()
	db.AutoMigrate(&User{})

	user := User{
		Email:    email,
		Password: password,
	}
	db.Create(&user)
}

func GetUserPassword(email string) string {
	postgresInfo := common.GetPostgresInfo()
	db, err := gorm.Open("postgres", postgresInfo)
	common.CheckErr(err)
	defer db.Close()
	db.AutoMigrate(&User{})

	var user User
	db.Where("email = ?", email).Find(&user)
	return user.Password
}

func GetAllUsers() Users {
	postgresInfo := common.GetPostgresInfo()
	db, err := gorm.Open("postgres", postgresInfo)
	common.CheckErr(err)
	defer db.Close()
	db.AutoMigrate(&User{})

	var users Users
	db.Find(&users)
	return users
}

func GetUserById(userId string) User {
	postgresInfo := common.GetPostgresInfo()
	db, err := gorm.Open("postgres", postgresInfo)
	common.CheckErr(err)
	defer db.Close()
	db.AutoMigrate(&User{})

	var user User
	db.Where("ID = ?", userId).Find(&user)
	return user
}
