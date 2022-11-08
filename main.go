package main

import (
	"log"
	"web_test/database"
	"web_test/redis"
	"web_test/webroute"

	"github.com/gin-gonic/gin"
)

func main() {

	var err error

	database.Connect() // DB 연결 및 전역선언

	redis.InitRedis() // redis 연결
	defer redis.CloseRedis()

	e := gin.Default() // gin 엔진 생성

	/*server := socketio.NewServer(nil)

	mywebsocket.Socket(server) // 소켓 셋팅
	defer server.Close()       // 소켓서버 닫기*/

	webroute.WebRoute(e /*, server*/) // gin 라우팅(소켓서버 포함)

	if err = e.Run(); err != nil { // http 서버열기 :8080
		log.Fatal("failed run server", err)
	}

	log.Println("서버 열기 성공")
}
