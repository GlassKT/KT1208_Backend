package controller

import (
	"io"
	"os"
	"web_test/database"
	"web_test/models"
	"web_test/modules"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func ImgUpload(c *gin.Context) {

	multiFile, header, err := c.Request.FormFile("image") // multi part 파일 받기
	if err != nil {
		response.ResponseBadRequest(c, "img upload dont find formfile(image)")
		return
	}

	fileName := header.Filename

	file, err := os.Create("./image/" + fileName) // 파일 생성
	if err != nil {
		response.ResponseBadRequest(c, "img upload create file error")
		return
	}

	res, err := io.Copy(file, multiFile) // 파일 쓰기
	if err != nil || res == 0 {
		response.ResponseBadRequest(c, "img upload copy file error")
		return
	}

	img_id := modules.FileName(fileName) // 이미지 파일에서 이름 뽑아오기

	if err = database.DB.Model(&models.User{}).Where("id = ?", img_id).Update("img_name", fileName).Error; err != nil { // db에 넣기
		response.ResponseBadRequest(c, "img upload update error")
		return
	}

	/*aiRes := modules.ImgAIapi(fileName) // 사람 얼굴인지 파악하는 api
	if !aiRes {
		response.ResponseBadRequest(c, "사람의 얼굴이 아닙니다")
		return
	}*/

	response.ResponseStatusAccept(c, "img upload success")
}
