package common

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "strings"
  jwt "github.com/dgrijalva/jwt-go"
)

func CheckAuth(c *gin.Context) bool {
  tokenValid := false
  if len(c.Request.Header.Get("Authorization")) > 0 {
	  requestToken := strings.TrimSpace(c.Request.Header.Get("Authorization")[7:len(c.Request.Header.Get("Authorization"))])

	  token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
	    return []byte("secret"), nil
	  })
	  checkErr(err)

    tokenValid = token.Valid
  }

  return tokenValid
}

func checkErr(err error) {
  if (err != nil) {
    fmt.Println("error = " + err.Error())
  }
}
