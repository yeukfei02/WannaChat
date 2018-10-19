package userController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"

	"WannaChat/common"
	"WannaChat/model/userModel"
	. "github.com/dgrijalva/jwt-go"
)

// request body
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	StandardClaims
}

func Signup(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	userModel.InsertUser(user.Email, user.Password)

	c.JSON(200, gin.H{
		"message":         "signup!",
		"user":            user,
		"createdDateTime": time.Now(),
	})
}

func Login(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	userPasswordFromDb := userModel.GetUserPassword(user.Email)

	if user.Password == userPasswordFromDb {
		token := NewWithClaims(GetSigningMethod("HS256"), &User{
			Email:    user.Email,
			Password: user.Password,
		})
		tokenString, err := token.SignedString([]byte("secret"))
		checkErr(err)

		c.JSON(200, gin.H{
			"message":         "login success!",
			"token":           tokenString,
			"createdDateTime": time.Now(),
		})
	} else {
		c.JSON(400, gin.H{
			"message": "login fail!",
		})
	}
}

func GetAllUsers(c *gin.Context) {
	tokenValid := common.CheckAuth(c)
	if tokenValid {
		usersList := userModel.GetAllUsers()

		c.JSON(200, gin.H{
			"message": "get all users!",
			"users":   usersList,
			"count":   len(usersList),
		})
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token!",
		})
	}
}

func GetUserById(c *gin.Context) {
	tokenValid := common.CheckAuth(c)
	if tokenValid {
		userId := c.Param("id")
		user := userModel.GetUserById(userId)

		c.JSON(200, gin.H{
			"message": "get user by id!",
			"user":    user,
		})
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token!",
		})
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("error = " + err.Error())
	}
}
