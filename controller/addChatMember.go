package controller

import (
	"web_test/database"
	"web_test/models"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func AddChattingMember(c *gin.Context) {

	chatRoom := &models.ChattingRoom{}

	err := c.Bind(chatRoom)
	if err != nil {
		response.ResponseBadRequest(c, "add chatting member bind error")
		return
	}

	if chatRoom.User == "" || chatRoom.RoomNum == "" {
		response.ResponseBadRequest(c, "유저이름과 방 번호가 비어있음")
		return
	}

	err = database.DB.Model(&models.ChattingRoom{}).Create(chatRoom).Error
	if err != nil {
		response.ResponseBadRequest(c, "add chatting member create error")
		return
	}

	response.ResponseStatusAccept(c, "add chatting member success")
}
