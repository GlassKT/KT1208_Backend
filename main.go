package main

import (
	"net/http"

	"example.com/m/database"

	"example.com/m/mysocket"
	"example.com/m/webRoute"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func main() {

	database.Connect() // DB 연결 및 전역선언
	e := gin.Default() // gin 엔진 생성

	webRoute.WRoute(e) // routing

	server := socketio.NewServer(nil) // socket 서버 생성

	mysocket.SRoute(server) // socket handling

	go server.Serve()    // socket server 열기
	defer server.Close() // socket server 닫기

	http.Handle("/socket.io/", server) // socket routing

	e.Run() // http 서버열기 :8080
}
