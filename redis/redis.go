package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client // 전역 변수
var dsn string = "localhost:6379"

func InitRedis() { // redis 연결

	Client = redis.NewClient(&redis.Options{ // redis 연결
		Addr: dsn,
	})

	if _, err := Client.Ping(context.Background()).Result(); err != nil { // redis 연결 확인
		panic(err)
	}

	log.Println("redis 연결 성공")
}

func CloseRedis() { // redis 연결 끊기
	Client.Close()
}
