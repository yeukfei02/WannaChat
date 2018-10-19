package userModel

import (
	"WannaChat/common"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Users []User

type User struct {
	gorm.Model
	UserId   int    `json:"userId" gorm:"primary_key; auto_increment;"`
	Email    string `json:"email" gorm:"type:varchar(255); not null;"`
	Password string `json:"password" gorm:"type:varchar(255); not null;"`
}

func InsertUser(email string, password string) {
	postgresInfo := common.GetPostgresInfo()
	db, err := gorm.Open("postgres", postgresInfo)
	checkErr(err)
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
	checkErr(err)
	defer db.Close()
	db.AutoMigrate(&User{})

	var user User
	db.Where("email = ?", email).Find(&user)
	return user.Password
}

func GetAllUsers() Users {
	postgresInfo := common.GetPostgresInfo()
	db, err := gorm.Open("postgres", postgresInfo)
	checkErr(err)
	defer db.Close()
	db.AutoMigrate(&User{})

	var users Users
	db.Find(&users)
	return users
}

func GetUserById(userId string) User {
	postgresInfo := common.GetPostgresInfo()
	db, err := gorm.Open("postgres", postgresInfo)
	checkErr(err)
	defer db.Close()
	db.AutoMigrate(&User{})

	var user User
	db.Where("ID = ?", userId).Find(&user)
	return user
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("error = " + err.Error())
	}
}
