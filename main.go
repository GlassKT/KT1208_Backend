package main

import (
	"GlassKT/database"
	"GlassKT/method"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	g.POST("/signup", method.SignUp(g.HandleContext()))
	g.Run()
}
