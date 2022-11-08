package controller

import (
	"web_test/database"
	"web_test/models"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func LeaveChatRoom(c *gin.Context) {

	chatRoom := &models.ChattingRoom{}

	err := c.Bind(chatRoom)
	if err != nil {
		response.ResponseBadRequest(c, "leave chat room bind error")
		return
	}

	if chatRoom.User == "" || chatRoom.RoomNum == "" {
		response.ResponseBadRequest(c, "아이디나 방 번호가 비어있습니다")
		return
	}

	err = database.DB.Model(&models.ChattingRoom{}).Where("user_id = ? and room_num = ?", chatRoom.User, chatRoom.RoomNum).Delete(chatRoom).Error
	if err != nil {
		response.ResponseBadRequest(c, "leave chat room delete error")
		return
	}

	response.ResponseStatusAccept(c, "leave chat room success")
}
