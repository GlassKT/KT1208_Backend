package controllers

import (
	db "example.com/m/database"
	"example.com/m/models"

	"github.com/gin-gonic/gin"
)

func AcceptFriend(c *gin.Context) {

	// 친구 수락 핸들러

	Makefriend := &models.Makefriend{}

	if err := c.Bind(Makefriend); err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"messege": "바인딩 오류",
		})
	}

	Makefriend.STATUS = "follow"

	if err := db.DB.Model(Makefriend).Omit("user_id", "friend_id").Where("user_id = ? and friend_id = ?", Makefriend.USERID, Makefriend.FRIENDID).Updates(Makefriend).Error; err != nil {
		c.JSON(400, gin.H{
			"status":  "400",
			"message": "id일때 업데이트에 실패함",
		})
		return
	}

	if err := db.DB.Model(Makefriend).Omit("user_id", "friend_id").Where("user_id = ? and friend_id = ?", Makefriend.FRIENDID, Makefriend.USERID).Updates(Makefriend).Error; err != nil {
		c.JSON(400, gin.H{
			"status":  "400",
			"message": "friendId일때 업데이트에 실패함",
		})
		return
	}
	return
}
