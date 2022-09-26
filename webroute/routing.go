package webroute

import (
	"log"
	"net/http"
	"web_test/controller"
	"web_test/middleware"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func WebRoute(e *gin.Engine, server *socketio.Server) {

	/*e.Static("/img", "./img")

	e.Use(middleware.CORSMiddleware())

	e.POST("/signup", controller.SignUp)                                                         // 회원가입 페이지
	e.POST("/login", controller.Login)                                                           // 로그인 페이지
	e.POST("/uploadimg", controller.ImgUpload)                                                   // 이미지 업로드
	e.POST("/logout", middleware.TokenAuthMiddleware(), controller.LogOut)                       // 로그 아웃
	e.POST("/edituser", middleware.TokenAuthMiddleware(), controller.EditUser)                   // 프로필 편집 페이지
	e.POST("/deletefriend", middleware.TokenAuthMiddleware(), controller.DeleteFriend)           // 친구 거절, 삭제
	e.POST("/deleteblockfriend", middleware.TokenAuthMiddleware(), controller.DeleteBlockFriend) // 친구 차단 삭제
	e.POST("/deletechatroom", middleware.TokenAuthMiddleware(), controller.DeleteChattingRoom)        // 채탕륨 석재
	e.POST("/acceptfriend", middleware.TokenAuthMiddleware(), controller.AcceptFriend)           // 친구 허용
	e.POST("/addfriend", middleware.TokenAuthMiddleware(), controller.AddFriend)                 // 친구 신청
	e.POST("/addhobby", middleware.TokenAuthMiddleware(), controller.AddHobby)                   // 취미 추가
	e.POST("/addblockfriend", middleware.TokenAuthMiddleware(), controller.AddBlockFriend)       // 친구 차단
	e.POST("/addchatroom", middleware.TokenAuthMiddleware(), controller.AddChattingRoom)         // 채팅룸 만들기

	e.POST("/", middleware.TokenAuthMiddleware()) // 로그인 확인

	e.GET("/emailauth", controller.EmailAuth)                                                             // 이메일 인증
	e.GET("/getwaitfriend", middleware.TokenAuthMiddleware(), controller.GetWaitFriend)                   // 친구 대기 리스트
	e.GET("/getfriend", middleware.TokenAuthMiddleware(), controller.GetFriend)                           // 친구 요청 리스트
	e.GET("/getuser", middleware.TokenAuthMiddleware(), controller.GetUser)                               // 유저 프로필 정보
	e.GET("/getblockfriend", middleware.TokenAuthMiddleware(), controller.GetBlockFriend)                 // 친구 차단 리스트
	e.GET("/recommandfriend", controller.RecommandFriend)                                                 // 랜덤 친구 추천
	e.GET("/recommandfriendfriends", middleware.TokenAuthMiddleware(), controller.RecommandFriendFriends) // 친구의 친구 추천

	log.Println("라우팅 성공")*/

	e.Static("/img", "./img")

	e.Use(middleware.CORSMiddleware())

	e.POST("/signup", controller.SignUp)                       // 회원가입 페이지
	e.POST("/login", controller.Login)                         // 로그인 페이지
	e.POST("/uploadimg", controller.ImgUpload)                 // 이미지 업로드
	e.POST("/logout", controller.LogOut)                       // 로그 아웃
	e.POST("/edituser", controller.EditUser)                   // 프로필 편집 페이지
	e.POST("/deletefriend", controller.DeleteFriend)           // 친구 거절, 삭제
	e.POST("/deleteblockfriend", controller.DeleteBlockFriend) // 친구 차단 삭제
	e.POST("/deletechatroom", controller.DeleteChattingRoom)   // 채탕륨 석재
	e.POST("/acceptfriend", controller.AcceptFriend)           // 친구 허용
	e.POST("/addfriend", controller.AddFriend)                 // 친구 신청
	e.POST("/addhobby", controller.AddHobby)                   // 취미 추가
	e.POST("/addblockfriend", controller.AddBlockFriend)       // 친구 차단
	e.POST("/addchatroom", controller.AddChattingRoom)         // 채팅룸 만들기

	e.POST("/", middleware.TokenAuthMiddleware()) // 로그인 확인

	e.GET("/emailauth", controller.EmailAuth)                           // 이메일 인증
	e.GET("/getwaitfriend", controller.GetWaitFriend)                   // 친구 대기 리스트
	e.GET("/getfriend", controller.GetFriend)                           // 친구 요청 리스트
	e.GET("/getuser", controller.GetUser)                               // 유저 프로필 정보
	e.GET("/getblockfriend", controller.GetBlockFriend)                 // 친구 차단 리스트
	e.GET("/recommandfriend", controller.RecommandFriend)               // 랜덤 친구 추천
	e.GET("/recommandfriendfriends", controller.RecommandFriendFriends) // 친구의 친구 추천

	// ------------  socket  ------------
	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error : %s\n", err)
		}
	}()

	e.GET("/socket.io/*any", gin.WrapH(server))
	e.POST("/socket.io/*any", gin.WrapH(server))
	e.StaticFS("/chat", http.Dir("./asset"))

	log.Println("라우팅 성공")

}
