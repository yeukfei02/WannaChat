package userController

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "strings"
	"time"

  "WannaChat/model/userModel"
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

  userModel.InsertUser(user.Email, user.Password)

	c.JSON(200, gin.H{
		"message": "signup!",
		"user": user,
		"createdDateTime": time.Now(),
	})
}

func Login(c *gin.Context) {
	var user User
	c.BindJSON(&user)

  userPasswordFromDb := userModel.GetUserPassword(user.Email)

  if (user.Password == userPasswordFromDb) {
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
  } else {
    c.JSON(400, gin.H{
      "message": "login fail!",
    })
  }
}

func GetAllUsers(c *gin.Context) {
  if len(c.Request.Header.Get("Authorization")) > 0 {
	  requestToken := strings.TrimSpace(c.Request.Header.Get("Authorization")[7:len(c.Request.Header.Get("Authorization"))])

	  token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
	    return []byte("secret"), nil
	  })
	  checkErr(err)

    if (token.Valid) {
			usersList := userModel.GetAllUsers()

		  c.JSON(200, gin.H{
		    "message": "get all users!",
		    "users": usersList,
		    "count": len(usersList),
		  })
		} else {
			c.JSON(404, gin.H{
		    "message": "wrong token!",
		  })
		}
  } else {
    c.JSON(404, gin.H{
			"message": "missing token!",
		})
  }
}

func GetUserById(c *gin.Context) {
  if len(c.Request.Header.Get("Authorization")) > 0 {
	  requestToken := strings.TrimSpace(c.Request.Header.Get("Authorization")[7:len(c.Request.Header.Get("Authorization"))])

	  token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
	    return []byte("secret"), nil
	  })
	  checkErr(err)

    if (token.Valid) {
      userId := c.Param("id")
			user := userModel.GetUserById(userId)

		  c.JSON(200, gin.H{
		    "message": "get user by id!",
		    "user": user,
		  })
		} else {
			c.JSON(404, gin.H{
		    "message": "wrong token!",
		  })
		}
  } else {
    c.JSON(404, gin.H{
			"message": "missing token!",
		})
  }
}

func checkErr(err error) {
  if (err != nil) {
    fmt.Println("error = " + err.Error())
  }
}
