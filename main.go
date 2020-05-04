package main

import (
	helmet "github.com/danielkov/gin-helmet"
	cors "github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"WannaChat/src/common"
	"WannaChat/src/routes/groupRoutes"
	"WannaChat/src/routes/mainRoutes"
	"WannaChat/src/routes/membershipRoutes"
	"WannaChat/src/routes/userRoutes"
	"WannaChat/src/schema"
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

	mainRoutes.Routes(r)
	userRoutes.Routes(r)
	groupRoutes.Routes(r)
	membershipRoutes.Routes(r)

	r.Run()
}
