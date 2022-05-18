package controllers

import (
	db "GlassKT/database"

	"github.com/gin-gonic/gin"
)

func GetHobby(g *gin.Context) {

	Hobby := []db.Hobby{} // db.hobby{}의 형태로 가져온 값들의 집합을 리스트에 저장

	if err := db.DB.Find(&Hobby).Error; err != nil {
		g.JSON(200, gin.H{
			"status": 200,
			"data":   "값을 가져오지 못함",
		})
		return
	}
	g.JSON(200, gin.H{
		"status": 200,
		"data":   Hobby,
	})
	return
}
