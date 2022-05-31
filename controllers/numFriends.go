package controllers

import (
	db "GlassKT/database"

	"github.com/gin-gonic/gin"
)

func NumFriends(g *gin.Context) {

	var cv int64
	Makefriend := &db.Makefriend{}

	g.Bind(Makefriend)

	if err := db.DB.Model(Makefriend).Where("user_id = ?", Makefriend.USERID).Count(&cv).Error; err != nil {
		g.JSON(400, gin.H{
			"status":  400,
			"message": "불러오지 못함",
		})
	}

	g.JSON(200, gin.H{
		"status":  200,
		"message": cv,
	})
	return
}
