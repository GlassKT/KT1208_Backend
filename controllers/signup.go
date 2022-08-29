package controllers

import (
	_ "database/sql"
	"log"

	"example.com/m/helper"

	db "example.com/m/database"
	"example.com/m/models"
	"example.com/m/modules"

	"github.com/gin-gonic/gin"
)

// 받은 데이터의 ID가 DB에 존재한다면 string, false 반환
// 없다면 받은 데이터를 데이터베이스에 넣고 string, True를 반환

func SignUp(c *gin.Context) {

	user := new(models.User) // 새로운 유저객체
	err := c.Bind(user)      // g의 값을 user에 파싱
	User := &models.User{}   // orm에서 원본 DB 접근을 위한 객체

	if err != nil {
		c.JSON(400, gin.H{"status": "400", "message": "파싱중 오류"})
		return
	}

	if user.ID == "" || user.PW == "" || user.NAME == "" {
		c.JSON(400, gin.H{"status": "400", "message": "ID,PW,NAME 모두 입력하세요"})
		return
	}

	// DB에서 id가 user.ID를 가진것이 data
	if err := db.DB.Where("id = ?", user.ID).Find(User).Error; err != nil {
		c.JSON(400, gin.H{"status": "400", "message": "이미 있는 ID"})
		return
	}

	user.PW = helper.HashPassword(user.PW)

	err = modules.MakeUser(user.ID, user.PW, user.NAME, user.EMAIL)
	if err != nil {
		c.JSON(400, gin.H{"status": "400", "message": "유저 생성중 오류"})
		return
	}
	c.JSON(200, gin.H{"status": "200", "message": "회원가입 성공"})

	log.Print("[회원가입 성공]")
	return
}
