package userController

import (
	"time"

	"github.com/badoux/checkmail"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	. "WannaChat/common"
	"WannaChat/model/userModel"
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

	if len(user.Email) > 0 && len(user.Password) > 0 {
		err := checkmail.ValidateFormat(user.Email)
		if err == nil {
			userModel.InsertUser(user.Email, user.Password)

			c.JSON(201, gin.H{
				"message":         "signup!",
				"user":            user,
				"createdDateTime": time.Now(),
			})
		} else {
			c.JSON(400, gin.H{
				"message": "invalid email format!",
			})
		}
	} else {
		c.JSON(400, gin.H{
			"message": "wrong request format!",
		})
	}
}

func Login(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	if len(user.Email) > 0 && len(user.Password) > 0 {
		err := checkmail.ValidateFormat(user.Email)
		if err == nil {
			userPasswordFromDb := userModel.GetUserPassword(user.Email)

			if user.Password == userPasswordFromDb {
				token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &User{
					Email:    user.Email,
					Password: user.Password,
				})
				tokenString, err := token.SignedString([]byte("secret"))
				CheckErr(err)

				c.JSON(201, gin.H{
					"message":         "login success!",
					"token":           tokenString,
					"createdDateTime": time.Now(),
				})
			} else {
				c.JSON(400, gin.H{
					"message": "login fail!",
				})
			}
		} else {
			c.JSON(400, gin.H{
				"message": "invalid email format!",
			})
		}
	} else {
		c.JSON(400, gin.H{
			"message": "wrong request format!",
		})
	}
}

func GetAllUsers(c *gin.Context) {
	tokenValid := CheckAuth(c)
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
	tokenValid := CheckAuth(c)
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
