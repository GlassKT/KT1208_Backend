package controllers

import (
	"math/rand"
	"strconv"
	"time"

	"example.com/m/models"
	"example.com/m/modules"

	"github.com/gin-gonic/gin"
)

func SendEmaill(c *gin.Context) {

	email := new(models.User)
	err := c.Bind(email)
	if err != nil {
		c.JSON(400, gin.H{"status": "400", "message": "파싱중 오류"})
		return
	}

	rand.Seed(time.Now().UnixNano())
	authNum := strconv.Itoa(rand.Intn(999999))

	if err := modules.SendMail(authNum, email.EMAIL); err != nil {
		c.JSON(400, gin.H{"status": "400", "message": "이메일 전송 실패"})
		return
	} else {
		c.JSON(400, gin.H{"status": "200", "message": "이메일 전송 성공", "authNum": authNum})
	}
	return
}
