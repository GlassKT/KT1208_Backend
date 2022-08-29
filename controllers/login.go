package controllers

import (
	db "example.com/m/database"
	"example.com/m/models"
	"example.com/m/modules"

	"github.com/gin-gonic/gin"
)

type loginstruct struct {
	ID string `json:"id"`
	PW string `json:"pw"`
}

func Login(c *gin.Context) {
	param := new(loginstruct)
	err := c.Bind(param) // g의 값을 user에 파싱
	if err != nil {
		c.JSON(400, gin.H{"status": "400", "message": "바인딩 중 오류"})
		return
	}

	User := &models.User{} // orm에서 원본 DB 접근을 위한 객체

	if param.ID == "" || param.PW == "" {
		c.JSON(400, gin.H{"status": "400", "message": "ID,PW 모두 입력하세요"})
		return
	}

	if err := db.DB.Where("id = ? AND pw = ?", param.ID, param.PW).Find(User).Error; err != nil {
		c.JSON(400, gin.H{"status": "400", "message": "ID가 존재하지 않거나 PW가 잘못되었습니다"})
		return
	}

	token, err := modules.CreateToken(param.ID)
	if err != nil {
		c.JSON(400, gin.H{"status": "400", "message": "토큰 생성중 오류"})
		return
	}
	c.SetCookie("access-token", token, 60, "/login", "localhost:8080", false, false)
	c.JSON(200, gin.H{"status": "200", "message": "토근생성", "accessToken": token})
	return
}
