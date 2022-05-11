package controllers

import (
	db "GlassKT/database"

	"github.com/gin-gonic/gin"
)

func Profile(g *gin.Context) {

	User := db.User{}
	//profile := proFile{}

	err := g.Bind(&User)

	if err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "바인딩 오류",
		})
		return
	}

	if err := db.DB.Where("id = ?", g.Param("id")).Find(&User).Error; err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "존재하지 않는 Id",
		})
		return
	}

	db.DB.Omit("pw", "createAt", "").First(&User, "id = ?", User.ID)

	g.JSON(200, gin.H{
		"status": "200",
		"data":   User,
	})

	return
}
