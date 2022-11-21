package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// bad

func ResponseBadRequest(c *gin.Context, msg string) { // 대부분의 실패시
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  http.StatusBadRequest,
		"message": msg,
	})
}

func ResponseUnauthorized(c *gin.Context, msg string) { // 토큰 인증 불가
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  http.StatusUnauthorized,
		"message": msg,
	})
}

// success

func ResponseStatusAccept(c *gin.Context, msg string) { // 대부분의 성공시
	c.JSON(http.StatusAccepted, gin.H{
		"status":  http.StatusAccepted,
		"message": msg,
	})
}

func ResponseToken(c *gin.Context, tk, id, name, image string) { // 로그인 토큰
	c.JSON(http.StatusAccepted, gin.H{
		"status":      http.StatusAccepted,
		"accessToken": tk,
		"id":          id,
		"name":        name,
		"image":       image,
	})
}

func ResponseAuthNum(c *gin.Context, authNum string) { // 이메일 인증번호
	c.JSON(http.StatusAccepted, gin.H{
		"status": http.StatusAccepted,
		"auth":   authNum,
	})
}

func ResponseList(c *gin.Context, f interface{}) {
	c.JSON(http.StatusAccepted, gin.H{
		"status": http.StatusAccepted,
		"data":   f,
	})
}
