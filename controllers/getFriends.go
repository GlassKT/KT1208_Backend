package controllers

import (
	db "GlassKT/database"

	"github.com/gin-gonic/gin"
)

func GetFriends(g *gin.Context) {

	User := []db.User{}

	userid := g.Query("userid")
	status := g.Query("status")

	// SELECT * FROM users WHERE id IN (SELECT friend_id FROM makefriends WHERE status = 'follow or wait' AND user_id = '아이디')
	subquery := db.DB.Table("makefriends").Select("friend_id").Where("status = ? and user_id = ?", status, userid)
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
