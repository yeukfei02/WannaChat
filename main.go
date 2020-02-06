package main

import (
	helmet "github.com/danielkov/gin-helmet"
	cors "github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"WannaChat/common"
	"WannaChat/routes/groupRoutes"
	"WannaChat/routes/membershipRoutes"
	"WannaChat/routes/userRoutes"
	"WannaChat/schema"
)

func connectDBAndCreateTable(table interface{}) {
	db := common.OpenPostgresDBLazy()
	defer db.Close()
	common.CheckTableExists(db, table)
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(helmet.Default())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	connectDBAndCreateTable(&schema.User{})
	connectDBAndCreateTable(&schema.Group{})
	connectDBAndCreateTable(&schema.Membership{})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "WannaChat api",
		})
	})

	userRoutes.Routes(r)
	groupRoutes.Routes(r)
	membershipRoutes.Routes(r)

	r.Run()
}
