package controllers

import (
	db "GlassKT/database"

	"github.com/gin-gonic/gin"
)

func DeleteFriends(g *gin.Context) {

	Makefriend := &db.Makefriend{}
	if err := g.Bind(Makefriend); err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "바인딩 오류",
		})
		return
	}

	if err := db.DB.Where("user_id = ? and friend_id = ?", Makefriend.USERID, Makefriend.FRIENDID).Delete(Makefriend).Error; err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "id일때 삭제에 실패함",
		})
		return
	}

	if err := db.DB.Where("user_id = ? and friend_id = ?", Makefriend.FRIENDID, Makefriend.USERID).Delete(Makefriend).Error; err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "friend_id일때 삭제에 실패함",
		})
		return
	}

	g.JSON(200, gin.H{
		"status":  "200",
		"message": "친구 삭제 완료",
	})
	return
}
