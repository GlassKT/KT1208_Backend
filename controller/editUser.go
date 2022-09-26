package controller

import (
	"web_test/database"
	"web_test/models"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func EditUser(c *gin.Context) {

	editInfo := &models.EditInfo{}

	err := c.Bind(editInfo) // bind
	if err != nil {
		response.ResponseBadRequest(c, "edit bind error")
		return
	}

	if err := database.DB.Where("id = ?", editInfo.Id).Find(&models.User{}).Error; err != nil { // id 존재여부 확인
		response.ResponseBadRequest(c, "edit id is not found")
		return
	}

	if err := database.DB.Model(&models.User{}).Omit("id", "pw").Where("id = ?", editInfo.Id).Updates(editInfo).Error; err != nil { // id, pw 제외 구조체 째로 업데이트
		response.ResponseBadRequest(c, "edit update error")
		return
	}

	response.ResponseStatusAccept(c, "edit success")
}
