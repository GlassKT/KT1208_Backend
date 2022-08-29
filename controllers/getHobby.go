package controllers

import (
	db "example.com/m/database"
	"example.com/m/models"

	"github.com/gin-gonic/gin"
)

func GetHobby(c *gin.Context) {

	Hobby := []models.Hobby{} // db.hobby{}의 형태로 가져온 값들의 집합을 리스트에 저장

	if err := db.DB.Find(&Hobby).Error; err != nil {
		c.JSON(200, gin.H{
			"status": 200,
			"data":   "값을 가져오지 못함",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"data":   Hobby,
	})
	return
}
