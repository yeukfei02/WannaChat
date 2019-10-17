package main

import (
	helmet "github.com/danielkov/gin-helmet"
	cors "github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"WannaChat/routes/groupRoutes"
	"WannaChat/routes/membershipRoutes"
	"WannaChat/routes/userRoutes"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(helmet.Default())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	userRoutes.Routes(r)
	groupRoutes.Routes(r)
	membershipRoutes.Routes(r)

	r.Run()
}
