package controllers

import (
	db "example.com/m/database"

	"example.com/m/models"

	"github.com/gin-gonic/gin"
)

func RecommandFriend(c *gin.Context) {

	User := []models.User{}

	result := map[string]interface{}{}

	if err := db.DB.Table("users").Where("id = ?", c.Query("id")).Take(&result).Error; err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"messege": "내 취미를 가지고 오는데 실패했습니다",
		})
		return
	}

	if err := db.DB.Select("id").Where("hobbyId = ? and id != ?", result["hobbyId"], result["id"]).Find(&User).Error; err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"messege": "취미가 같은 친구를 데려오는데 실패했습니다",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": 200,
		"data":   User,
	})
	return
}
