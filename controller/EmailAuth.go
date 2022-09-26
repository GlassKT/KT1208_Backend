package controller

// emailAuth로 할 경우 대소문자 구분의 어색한 golang씨발 때문에 오류남 gopls build tag 인가 뭔가 해서 바꿔야댐

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"web_test/modules"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func EmailAuth(c *gin.Context) {

	email := c.Query("email") // get query
	if email == "" {
		response.ResponseBadRequest(c, "email auth get query error")
	}

	fmt.Println(email)

	rand.Seed(time.Now().UnixNano())
	authNum := strconv.Itoa(rand.Intn(999999)) // 인증번호 랜덤으로 생성

	if err := modules.SendMail(authNum, email); err != nil { // send mail
		response.ResponseBadRequest(c, "email auth send mail error")
		return
	}

	response.ResponseAuthNum(c, authNum)
}
