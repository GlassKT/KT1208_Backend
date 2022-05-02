package controllers

import (
	db "GlassKT/database"
	"GlassKT/models"

	"github.com/gin-gonic/gin"
)

type loginstruct struct {
	ID string `json:"id"`
	PW string `json:"pw"`
}

func Login(g *gin.Context) {
	param := new(loginstruct)
	err := g.Bind(param) // g의 값을 user에 파싱
	if err != nil {
		g.JSON(400, gin.H{"status": "400", "message": "파싱중 오류"})
		return
	}

	User := &db.User{} // orm에서 원본 DB 접근을 위한 객체

	if param.ID == "" || param.PW == "" {
		g.JSON(400, gin.H{"status": "400", "message": "ID,PW 모두 입력하세요"})
		return
	}

	if err := db.DB.Where("id = ? AND pw = ?", param.ID, param.PW).Find(User).Error; err == nil {
		g.JSON(400, gin.H{"status": "400", "message": "ID가 존재하지 않거나 PW가 잘못되었습니다"})
		return
	}

	token, err := models.CreateToken(param.ID)
	if err != nil {
		g.JSON(400, gin.H{"status": "400", "message": "토큰 생성중 오류"})
		return
	}
	g.SetCookie("access-token", token, 60, "/login", "localhost:8080", false, false)
	g.JSON(200, gin.H{"status": "200", "message": "토근생성", "accessToken": token})
	return
}
