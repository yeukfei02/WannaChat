package common

import (
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	host       = "ec2-184-73-197-211.compute-1.amazonaws.com"
	port       = "5432"
	user       = "nfsqmmqiirrfxf"
	dbName     = "d9qd4thbsdmqkp"
	dbPassword = "0f4a0aa4b34a48fd5586772b743de5abeac903bec98ce98e44c1ca2bd6a7ac07"
)

func GetPostgresInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbName, dbPassword)
}

func CheckAuth(c *gin.Context) bool {
	tokenValid := false
	if len(c.Request.Header.Get("Authorization")) > 0 {
		requestToken := strings.TrimSpace(c.Request.Header.Get("Authorization")[7:len(c.Request.Header.Get("Authorization"))])

		token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		CheckErr(err)

		tokenValid = token.Valid
	}

	return tokenValid
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println("error = " + err.Error())
	}
}
