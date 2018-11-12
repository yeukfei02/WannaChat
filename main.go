package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

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
