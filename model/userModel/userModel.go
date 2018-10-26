package userModel

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"WannaChat/common"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"type:varchar(255); unique; not null;"`
	Password string `json:"password" gorm:"type:varchar(255); not null;"`
}

func InsertUser(email string, password string) {
	db, err := common.OpenPostgresDB()
	common.CheckErr(err)
	defer db.Close()
	common.CheckTableExists(db, &User{})

	user := User{
		Email:    email,
		Password: password,
	}
	db.Create(&user)
}

func GetUserPassword(email string) string {
	db, err := common.OpenPostgresDB()
	common.CheckErr(err)
	defer db.Close()
	common.CheckTableExists(db, &User{})

	var user User
	db.Where("email = ?", email).Find(&user)
	return user.Password
}

func GetAllUsers() []User {
	db, err := common.OpenPostgresDB()
	common.CheckErr(err)
	defer db.Close()
	common.CheckTableExists(db, &User{})

	var users []User
	db.Find(&users)
	return users
}

func GetUserById(userId string) User {
	db := common.OpenPostgresDBLazy()
	defer db.Close()
	common.CheckTableExists(db, &User{})

	var user User
	db.Where("ID = ?", userId).Find(&user)
	return user
}
