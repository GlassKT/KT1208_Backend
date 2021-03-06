package controllers

import (
	db "GlassKT/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 클라이언트에서 프랜드 유저의 id를 보내주면 프렌드테이블에 저장 // status도 같이 전달받음

func MakeFriend(g *gin.Context) {

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

	if err = db.DB.Where("id = ?", Makefriend.FRIENDID).Find(User).Error; err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "존재하지 않는 id",
		})
		return
	}

	Makefriend.STATUS = "wait" // 기본 값

	if err = db.DB.Create(Makefriend).Error; err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "인설트 오류",
		})
		return
	}

	temp := Makefriend.USERID
	Makefriend.USERID = Makefriend.FRIENDID
	Makefriend.FRIENDID = temp

	if err = db.DB.Create(Makefriend).Error; err != nil {
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
