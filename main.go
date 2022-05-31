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

	g.Static("/img", "./img")

	g.POST("/signup", controllers.SignUp)        // 회원가입 페이지
	g.POST("/mailsend", controllers.SendEmaill)  // 메일을 보낼 주소를 받는 주소
	g.POST("/login", controllers.Login)          // 로그인 페이지
	g.POST("/upload", controllers.UploadIMG)     // 이미지 업로드 페이지
	g.POST("/friend", controllers.MakeFriend)    // 친구추가하는 페이진데 후에 models로 좌천될 예정
	g.POST("/edit", controllers.EditUser)        // 프로필 편집 페이지
	g.POST("/profile", controllers.Profile)      // 프로필 보여주는 페이지
	g.POST("/numfriend", controllers.NumFriends) // 친구수를 보내주는 페이지

	g.GET("/hobby", controllers.GetHobby) // 취미목록을 보내주는 페이지

	g.Run() // :8080
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
