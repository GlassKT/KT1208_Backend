package controller

import (
	"web_test/database"
	"web_test/models"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func DeleteChattingRoom(c *gin.Context) {

	chatRoom := &models.ChattingRoom{}

	err := c.Bind(chatRoom)
	if err != nil {
		response.ResponseBadRequest(c, "delete chatting room bind error")
		return
	}

	if chatRoom.RoomNum == "" {
		response.ResponseBadRequest(c, "방 번호가 비어있음")
		return
	}

	if err = database.DB.Delete(chatRoom).Error; err != nil {
		response.ResponseBadRequest(c, "delete chatting room create error")
		return
	}

	response.ResponseStatusAccept(c, "delete chatting room success")

}
