package userController

import (
  "fmt"
  "github.com/gin-gonic/gin"
	"time"

  // "WannaChat/model/userModel"
  jwt "github.com/dgrijalva/jwt-go"
)

// request body
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func Signup(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	c.JSON(200, gin.H{
		"message": "signup!",
		"user": user,
		"createdDateTime": time.Now(),
	})
}

func Login(c *gin.Context) {
	var user User
	c.BindJSON(&user)

  token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &User{
    Email: user.Email,
    Password: user.Password,
  })
  tokenString, err := token.SignedString([]byte("secret"))
  checkErr(err)

  c.JSON(200, gin.H{
    "message": "login success!",
    "token": tokenString,
    "createdDateTime": time.Now(),
  })
}

func checkErr(err error) {
  if (err != nil) {
    fmt.Println("error = " + err.Error())
  }
}
