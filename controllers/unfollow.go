package controllers

import (
	"fmt"

	db "example.com/m/database"
	"example.com/m/models"

	"github.com/gin-gonic/gin"
)

func Unfollow(c *gin.Context) {

	User := &models.User{}
	Makefriend := &models.Makefriend{}
	err := c.Bind(Makefriend)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":  "400",
			"message": "바인딩 오류",
		})
		return
	}

	if err = db.DB.Where("id = ?", Makefriend.USERID).Find(User).Error; err != nil {
		c.JSON(400, gin.H{
			"status":  "400",
			"message": "존재하지 않는 id",
		})
		return
	}

	err = db.DB.Create(Makefriend).Error

	if err != nil {
		c.JSON(400, gin.H{
			"status":  "400",
			"message": "인설트 오류",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "200",
		"message": "친구 추가 완료",
	})
	return
}
