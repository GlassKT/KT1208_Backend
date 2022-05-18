package controllers

import (
	db "GlassKT/database"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func GetHobby(g *gin.Context) {

	Hobby := &db.Hobby{}
	//hobby := db.DB.Find(Hobby)

	/*if Hobby == nil {
		g.JSON(400, gin.H{
			"status":  400,
			"message": "바인딩 중 오류",
		})
		return
	}
	fmt.Println(hobby.RowsAffected)
	g.JSON(200, gin.H{
		"status":  400,
		"message": "취미목록",
		//"data":    hobby,
	})*/

	first := db.DB.First(Hobby)
	fmt.Println("firstValue", first.RowsAffected)
	many := db.DB.Find(Hobby)
	fmt.Println("many : ", many.RowsAffected)
	if err := db.DB.First(Hobby).Error; err != nil {
		log.Panic(err)
	}
	return
}
