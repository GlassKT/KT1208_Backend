package controllers

import (
	db "GlassKT/database"
	"GlassKT/models"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadIMG(g *gin.Context) {
	User := &db.User{}

	file, header, err := g.Request.FormFile("file")
	if err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "멀티파트폼 받기 실패",
		})
		return
	}

	imgId := g.Request.PostForm.Get("id")
	if imgId == "" {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "id받기 실패",
		})
		return
	}

	if err = db.DB.Where("id = ?", imgId).Find(User).Error; err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "존재하지 않는 id",
		})
		return
	}

	filename := header.Filename

	ext, err := models.FilePath(filename) // 확장자 뽑아내기
	if err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "파일 확장자를 못찾음",
		})
		return
	}

	imgIdOnly := imgId + ext // 파일명(유저 id) + 확장자명

	out, err := os.Create("./img/" + imgIdOnly)
	if err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "이미지 파일만들기 실패",
		})
		return
	}

	_, err = io.Copy(out, file)
	if err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "이미지 업로드",
		})
		return
	}

	picture := map[string]string{
		// "image": "http://10.80.162.125:8080/img/data.png",
		"image": "http://10.80.162.125:8080/img/" + imgIdOnly,
	}
	// }

	imgJson, err := json.Marshal(picture)
	if err != nil {
		g.JSON(400, gin.H{
			"status":  400,
			"messege": "map to json error",
		})
		return
	}

	buff := bytes.NewBuffer(imgJson)

	resp, err := http.Post("http://10.80.161.150:5000/image", "application/json", buff)
	if err != nil {
		panic(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	println(respBody)
	if err == nil {
		str := string(respBody)
		println(str)
	}

	// if err != nil {
	// 	g.JSON(400, gin.H{
	// 		"status":  400,
	// 		"messege": "Server to Server Post error",
	// 	})
	// 	return
	// }

	respResult := string(respBody)
	g.JSON(400, gin.H{
		"status":  400,
		"messege": respResult,
	})

	defer resp.Body.Close()
	defer out.Close()

	if err = db.DB.Model(&User).Where("id = ?", imgId).Update("imgname", imgIdOnly).Error; err != nil {
		g.JSON(400, gin.H{
			"status":  "400",
			"message": "이미지 업로드 실패",
		})
		return
	}

	g.JSON(200, gin.H{
		"status":  "200",
		"message": "이미지 업로드 성공",
	})
	return
}
