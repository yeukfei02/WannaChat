package main

import (
	helmet "github.com/danielkov/gin-helmet"
	cors "github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"WannaChat/src/helpers"
	"WannaChat/src/routes/groupRoutes"
	"WannaChat/src/routes/mainRoutes"
	"WannaChat/src/routes/membershipRoutes"
	"WannaChat/src/routes/userRoutes"

	"WannaChat/src/model/groupModel"
	"WannaChat/src/model/membershipModel"
	"WannaChat/src/model/userModel"
)

func connectDBAndCreateTable(table interface{}) {
	db := helpers.OpenPostgresDBLazy()
	defer db.Close()
	helpers.CheckTableExists(db, table)
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(helmet.Default())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	connectDBAndCreateTable(&userModel.User{})
	connectDBAndCreateTable(&groupModel.Group{})
	connectDBAndCreateTable(&membershipModel.Membership{})

	mainRoutes.Routes(r)
	userRoutes.Routes(r)
	groupRoutes.Routes(r)
	membershipRoutes.Routes(r)

	r.Run()
}
