package main

import (
	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"WannaChat/routes/groupRoutes"
	"WannaChat/routes/membershipRoutes"
	"WannaChat/routes/userRoutes"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	userRoutes.Routes(r)
	groupRoutes.Routes(r)
	membershipRoutes.Routes(r)

	r.Run()
}
