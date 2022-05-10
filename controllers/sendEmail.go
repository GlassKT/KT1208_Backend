package controllers

import (
	"GlassKT/database"
	"GlassKT/models"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SendEmaill(g *gin.Context) {

	email := new(database.User)
	err := g.Bind(email)
	if err != nil {
		g.JSON(400, gin.H{"status": "400", "message": "파싱중 오류"})
		return
	}

	rand.Seed(time.Now().UnixNano())
	authNum := strconv.Itoa(rand.Intn(999999))

	if err := models.SendMail(authNum, email.EMAIL); err != nil {
		g.JSON(400, gin.H{"status": "400", "message": "이메일 전송 실패"})
		return
	} else {
		g.JSON(400, gin.H{"status": "200", "message": "이메일 전송 성공", "authNum": authNum})
	}
	return
}
