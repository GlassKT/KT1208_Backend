package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// bad

func ResponseBadRequest(c *gin.Context, msg string) { // 대부분의 실패시
	c.JSON(http.StatusBadRequest, gin.H{"message": msg})
}

func ResponseUnauthorized(c *gin.Context, msg string) { // 토큰 인증 불가
	c.JSON(http.StatusUnauthorized, gin.H{"message": msg})
}

// success

func ResponseStatusAccept(c *gin.Context, msg string) { // 대부분의 성공시
	c.JSON(http.StatusAccepted, gin.H{"message": msg})
}

func ResponseToken(c *gin.Context, tk, id string) { // 로그인 토큰
	c.JSON(http.StatusAccepted, gin.H{"accessToken": tk, "id": id})
}

func ResponseAuthNum(c *gin.Context, authNum string) { // 이메일 인증번호
	c.JSON(http.StatusAccepted, gin.H{"auth": authNum})
}

func ResponseList(c *gin.Context, f interface{}) {
	c.JSON(http.StatusAccepted, gin.H{"data": f})
}
