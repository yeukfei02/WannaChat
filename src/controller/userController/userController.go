package userController

import (
	"os"

	"github.com/badoux/checkmail"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	"WannaChat/src/helpers"
	"WannaChat/src/services/userService"
)

// User request body
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// Signup controller
func Signup(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	if len(user.Email) > 0 && len(user.Password) > 0 {
		err := checkmail.ValidateFormat(user.Email)
		if err == nil {
			userPasswordByte := []byte(user.Password)
			hashedPassword, err := bcrypt.GenerateFromPassword(userPasswordByte, bcrypt.DefaultCost)
			helpers.CheckErr(err)
			userService.InsertUser(user.Email, string(hashedPassword))

			c.JSON(201, gin.H{
				"message": "signup success",
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

// Login controller
func Login(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	if len(user.Email) > 0 && len(user.Password) > 0 {
		err := checkmail.ValidateFormat(user.Email)
		if err == nil {
			userHashedPasswordFromDb := userService.GetUserPassword(user.Email)
			userHashedPasswordFromDbByte := []byte(userHashedPasswordFromDb)
			userPasswordByte := []byte(user.Password)

			passwordErr := bcrypt.CompareHashAndPassword(userHashedPasswordFromDbByte, userPasswordByte)
			if passwordErr == nil {
				token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &User{
					Email:    user.Email,
					Password: user.Password,
				})

				err := godotenv.Load()
				helpers.CheckErr(err)
				jwtSecret := os.Getenv("JWT_SECRET")

				tokenString, err := token.SignedString([]byte(jwtSecret))
				helpers.CheckErr(err)

				c.JSON(201, gin.H{
					"message": "login success",
					"token":   tokenString,
				})
			} else {
				c.JSON(400, gin.H{
					"message": "login fail, wrong password",
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

// GetAllUsers controller
func GetAllUsers(c *gin.Context) {
	tokenValid := helpers.CheckAuth(c)
	if tokenValid {
		usersList := userService.GetAllUsers()

		c.JSON(200, gin.H{
			"message": "get all users",
			"users":   usersList,
		})
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token",
		})
	}
}

// GetUserByID controller
func GetUserByID(c *gin.Context) {
	tokenValid := helpers.CheckAuth(c)
	if tokenValid {
		userID := c.Param("id")
		user := userService.GetUserByID(userID)

		c.JSON(200, gin.H{
			"message": "get user by id",
			"user":    user,
		})
	} else {
		c.JSON(404, gin.H{
			"message": "wrong or missing token",
		})
	}
}
