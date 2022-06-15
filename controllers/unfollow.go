package controllers

import (
	db "GlassKT/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Unfollow(g *gin.Context) {

	User := &db.User{}
	Makefriend := &db.Makefriend{}
	err := g.Bind(Makefriend)

	if err != nil {
		fmt.Println(err)
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "바인딩 오류",
		})
		return
	}

	if err = db.DB.Where("id = ?", Makefriend.USERID).Find(User).Error; err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "존재하지 않는 id",
		})
		return
	}

	err = db.DB.Create(Makefriend).Error

	if err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "인설트 오류",
		})
		return
	}

	g.JSON(200, gin.H{
		"status":  "200",
		"message": "친구 추가 완료",
	})
	return
}
