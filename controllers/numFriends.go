package controllers

import (
	db "example.com/m/database"
	"example.com/m/models"

	"github.com/gin-gonic/gin"
)

func NumFriends(c *gin.Context) {

	var cv int64
	Makefriend := &models.Makefriend{}

	c.Bind(Makefriend)

	if err := db.DB.Model(Makefriend).Where("user_id = ?", Makefriend.USERID).Count(&cv).Error; err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "불러오지 못함",
		})
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": cv,
	})
	return
}
