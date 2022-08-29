package controllers

import (
	db "example.com/m/database"
	"example.com/m/models"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {

	User := models.User{}
	//profile := proFile{}

	err := c.Bind(&User)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "400",
			"message": "바인딩 오류",
		})
		return
	}

	if err := db.DB.Where("id = ?", c.Param("id")).Find(&User).Error; err != nil {
		c.JSON(400, gin.H{
			"status":  "400",
			"message": "존재하지 않는 Id",
		})
		return
	}

	db.DB.Omit("pw", "createAt", "").First(&User, "id = ?", User.ID)

	c.JSON(200, gin.H{
		"status": "200",
		"data":   User,
	})

	return
}
