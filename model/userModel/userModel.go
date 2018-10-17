package userModel

import (
  "fmt"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Users []User

type User struct {
  gorm.Model
  UserId    int     `json:"userId" gorm:"primary_key; auto_increment;"`
	Email     string  `json:"email" gorm:"type:varchar(255); not null;"`
	Password  string  `json:"password" gorm:"type:varchar(255); not null;"`
}

const (
  host = "ec2-184-73-197-211.compute-1.amazonaws.com"
  port = "5432"
  user = "nfsqmmqiirrfxf"
  dbname = "d9qd4thbsdmqkp"
  dbPassword = "0f4a0aa4b34a48fd5586772b743de5abeac903bec98ce98e44c1ca2bd6a7ac07"
)

func InsertUser(email string, password string) {
  db, err := gorm.Open("postgres", "host=" + host + " port=" + port + " user=" + user + " dbname=" + dbname + " password=" + dbPassword)
  checkErr(err)
  defer db.Close()
  db.AutoMigrate(&User{})

  user := User{
    Email: email,
    Password: password,
  }
  db.Create(&user)
}

func GetUserPassword(email string) string {
  db, err := gorm.Open("postgres", "host=" + host + " port=" + port + " user=" + user + " dbname=" + dbname + " password=" + dbPassword)
  checkErr(err)
  defer db.Close()
  db.AutoMigrate(&User{})

  var user User
  db.Where("email = ?", email).Find(&user)
  return user.Password
}

func GetAllUsers() Users {
  db, err := gorm.Open("postgres", "host=" + host + " port=" + port + " user=" + user + " dbname=" + dbname + " password=" + dbPassword)
  checkErr(err)
  defer db.Close()
  db.AutoMigrate(&User{})

  var users Users
  db.Find(&users)
  return users
}

func GetUserById(userId string) User {
  db, err := gorm.Open("postgres", "host=" + host + " port=" + port + " user=" + user + " dbname=" + dbname + " password=" + dbPassword)
  checkErr(err)
  defer db.Close()
  db.AutoMigrate(&User{})

  var user User
  db.Where("ID = ?", userId).Find(&user)
  return user
}

func checkErr(err error) {
  if (err != nil) {
    fmt.Println("error = " + err.Error())
  }
}