package userModel

import (
  "fmt"

  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
  gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
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

  user := User{
    Email: email,
    Password: password,
  }
  db.Create(&user)
}

func GetUserPassword(email string) {
  db, err := gorm.Open("postgres", "host=" + host + " port=" + port + " user=" + user + " dbname=" + dbname + " password=" + dbPassword)
  checkErr(err)
  defer db.Close()

  user := User{
    Email: email,
  }
  db.Where("email = ?", email).Find(&user)
}

func checkErr(err error) {
  if (err != nil) {
    fmt.Println("error = " + err.Error())
  }
}
