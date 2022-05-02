package main

import (
	"GlassKT/controllers"
	"GlassKT/database"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()
	g := gin.Default()

	g.Use(CORSMiddleware())

	g.POST("/signup", controllers.SignUp)
	g.POST("/login", controllers.Login)
	g.POST("/edit", controllers.EditUser)
	g.POST("/upload", controllers.UploadIMG)
	g.POST("/friend", controllers.MakeFriend)
	//g.POST("/makefriend", controllers.MakeFriend)

	g.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
