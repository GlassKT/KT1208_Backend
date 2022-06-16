package controllers

import (
	db "GlassKT/database"

	"github.com/gin-gonic/gin"
)

func GetRequestFriends(g *gin.Context) {

	// 친구 목록 (내가 요청한 친구)
	User := []db.User{}

	userid := g.Query("userid")

	// SELECT * FROM users WHERE id IN (SELECT user_id FROM makefriends WHERE status = 'wait' AND friend_id = '아이디')
	subquery := db.DB.Table("makefriends").Select("user_id").Where("status = 'wait' and friend_id = ?", userid)
	query := db.DB.Omit("pw", "createAt").Where("id IN (?)", subquery).Find(&User)

	if err := query.Error; err != nil {
		g.JSON(200, gin.H{
			"status": 200,
			"data":   "값을 가져오지 못함",
		})
		return
	}

	g.JSON(200, gin.H{
		"status": 200,
		"data":   User,
	})
	return
}
