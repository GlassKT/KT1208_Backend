package controllers

import (
	"GlassKT/database"
	"GlassKT/models"

	"github.com/gin-gonic/gin"
)

func UploadIMG(g *gin.Context) {
	img := &database.User{}

	err := g.Bind(img)
	if err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "바인딩 오류",
		})
		return
	}

	err = models.StoreImg(img.ID, img.IMGNAME)
	if err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "이미지 업로드 오류",
		})
		return
	}
	g.JSON(200, gin.H{
		"status":  "200",
		"message": "이미지 업로드 성공",
	})
	return
}
