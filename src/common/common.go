package common

import (
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getPostgresInfo() string {
	err := godotenv.Load()
	CheckErr(err)
	host := os.Getenv("HOST")
	port := os.Getenv("PORT_NUMBER")
	user := os.Getenv("USERNAME")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbName, dbPassword)
}

// OpenPostgresDB common func
func OpenPostgresDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", getPostgresInfo())
	return db, err
}

// OpenPostgresDBLazy common func
func OpenPostgresDBLazy() *gorm.DB {
	db, err := gorm.Open("postgres", getPostgresInfo())
	CheckErr(err)
	return db
}

// CheckTableExists common func
func CheckTableExists(db *gorm.DB, table interface{}) {
	if !db.HasTable(table) {
		db.AutoMigrate(table)
	}
}

// CheckAuth common func
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

// CheckErr common func
func CheckErr(err error) {
	if err != nil {
		fmt.Println("error = " + err.Error())
	}
}
