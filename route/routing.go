package webRoute

import (
	"example.com/m/controllers"
	"example.com/m/middleware"
	"github.com/gin-gonic/gin"
)

func WRoute(e *gin.Engine) {

	e.Use(middleware.CORSMiddleware())

	e.Static("/img", "./img")

	e.POST("/signup", controllers.SignUp)              // 회원가입 페이지
	e.POST("/mailsend", controllers.SendEmaill)        // 메일을 보낼 주소를 받는 주소
	e.POST("/login", controllers.Login)                // 로그인 페이지
	e.POST("/upload", controllers.UploadIMG)           // 이미지 업로드 페이지
	e.POST("/makefriend", controllers.MakeFriend)      // 친구추가하는 페이진데 후에 models로 좌천될 예정
	e.POST("/edit", controllers.EditUser)              // 프로필 편집 페이지
	e.POST("/profile", controllers.Profile)            // 프로필 보여주는 페이지
	e.POST("/numfriend", controllers.NumFriends)       // 친구수를 보내주는 페이지
	e.POST("/unfollow", controllers.Unfollow)          // 친구 삭제
	e.POST("/deletefriend", controllers.DeleteFriends) // 친구 삭제
	e.POST("/acceptfriend", controllers.AcceptFriend)  // 친구 추가

	e.GET("/getrequestlist", controllers.GetRequestFriends)
	e.GET("/recommendFriend", controllers.RecommandFriend) // 취미가 같은 친구를 꺼내 보여주는 페이지
	e.GET("/hobby", controllers.GetHobby)                  // 취미목록을 보내주는 페이지
	e.GET("/friendList", controllers.GetFriends)           // 친구(서로친구 or 요청받은 친구 목록) 조회(아이디, 상태[wait, follow])
}
