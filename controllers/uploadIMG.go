package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	db "example.com/m/database"
	"example.com/m/models"
	"example.com/m/modules"

	"github.com/gin-gonic/gin"
)

func UploadIMG(c *gin.Context) {
	User := &models.User{}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "400",
			"message": "멀티파트폼 받기 실패",
		})
		return
	}

	imgId := c.Request.PostForm.Get("id")
	if imgId == "" {
		c.JSON(400, gin.H{
			"status":  "400",
			"message": "id받기 실패",
		})
		return
	}

	if err = db.DB.Where("id = ?", imgId).Find(User).Error; err != nil {
		c.JSON(400, gin.H{
			"status":  "400",
			"message": "존재하지 않는 id",
		})
		return
	}

	filename := header.Filename

	ext, err := modules.FilePath(filename) // 확장자 뽑아내기
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "400",
			"message": "파일 확장자를 못찾음",
		})
		return
	}

	imgIdOnly := imgId + ext // 파일명(유저 id) + 확장자명

	out, err := os.Create("./img/" + imgIdOnly)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "400",
			"message": "이미지 파일만들기 실패",
		})
		return
	}

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "400",
			"message": "이미지 업로드",
		})
		return
	}

	picture := map[string]string{
		// "image": "http://10.80.162.125:8080/img/data.png",
		"image": "http://10.80.162.33:8080/img/" + imgIdOnly, //***********
	}
	// }

	imgJson, err := json.Marshal(picture)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"messege": "map to json error",
		})
		return
	}

	buff := bytes.NewBuffer(imgJson)

	resp, err := http.Post("http://10.80.161.150:5000/image", "application/json", buff) //***********
	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"messege": "얼굴인식 서버와 연결이 되지 않음",
		})
		return
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"messege": "얼굴인식 판단 결과를 가져오지 못함",
		})
		return
	}

	respResult := string(respBody)

	if strings.Contains(respResult, "false") {
		c.JSON(400, gin.H{
			"status":  400,
			"messege": "사람의 얼굴이 없습니다. 얼굴사진을 올려주세요",
		})
		return
	}
	if strings.Contains(respResult, "error") {
		c.JSON(400, gin.H{
			"status":  400,
			"messege": "얼굴인식 서버에서 문제가 발생했습니다",
		})
		return
	}

	defer resp.Body.Close()
	defer out.Close()

	if err = db.DB.Model(&User).Where("id = ?", imgId).Update("imgname", imgIdOnly).Error; err != nil {
		c.JSON(400, gin.H{
			"status":  "400",
			"message": "이미지 업로드 실패",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "200",
		"message": "이미지 업로드 성공",
	})
	return
}
