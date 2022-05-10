package controllers

import (
	db "GlassKT/database"

	"github.com/gin-gonic/gin"
)

func Profile(g *gin.Context) {

	User := db.User{}
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

	db.DB.First(&User, "id = ?", User.ID)

	/*if User == nil{
		g.JSON(400,gin.H{
			"status":
		})
	}*/

	return
}
