package controllers

import (
	db "GlassKT/database"

	"github.com/gin-gonic/gin"
)

func EditUser(g *gin.Context) {
	newUser := new(db.User)
	err := g.Bind(newUser)
	if err != nil {
		g.JSON(400, gin.H{"status": "400", "message": "바인딩오류"})
		return
	}

	User := &db.User{}

	if newUser.ID == "" || newUser.PW == "" || newUser.NAME == "" {
		g.JSON(400, gin.H{"status": "400", "message": "ID,PW,NAME 모두 입력하세요"})
		return
	}

	if err := db.DB.Where("id = ? AND pw = ?", newUser.ID, newUser.PW).Find(User).Error; err == nil {
		g.JSON(400, gin.H{"status": "400", "message": "ID가 존재하지 않거나 PW가 잘못되었습니다"})
		return
	}

	if err := db.DB.Model(&User).Omit("createAt").Updates(newUser).Where("id = ?", newUser.ID).Error; err != nil {
		g.JSON(400, gin.H{"status": "400", "message": "업데이트 실패"})
		return
	}

	g.JSON(200, gin.H{"status": "200", "message": "업데이트 성공"})
	return
}
