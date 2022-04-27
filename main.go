package main

import (
	"GlassKT/database"
	"GlassKT/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()
	g := gin.Default()

	g.POST("/signup", handler.SignUp)
	g.POST("/login", handler.Login)

	g.Run()
}
