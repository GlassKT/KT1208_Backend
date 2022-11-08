package controller

import (
	"web_test/database"
	"web_test/models"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func GetRoomNum(c *gin.Context) {

	myRoom := []models.ChattingRoom{}
	friendRoom := []models.ChattingRoom{}

	roomnum := ""

	myId := c.Query("user")
	myFriend := c.Query("friend")

	if myFriend == "" || myId == "" {
		response.ResponseBadRequest(c, "게스트나 유저의 아이디가 비어있음")
		return
	}

	err := database.DB.Model(&models.ChattingRoom{}).Where("user = ?", myId).Find(&myRoom).Error
	if err != nil {
		response.ResponseBadRequest(c, "get room num find first error")
		return
	}

	err = database.DB.Model(&models.ChattingRoom{}).Where("user = ?", myFriend).Find(&friendRoom).Error
	if err != nil {
		response.ResponseBadRequest(c, "get room num find second error")
		return
	}

	for _, v := range myRoom {
		for _, v2 := range friendRoom {
			if v.RoomNum == v2.RoomNum {
				roomnum = v.RoomNum
				break
			}
		}
	}

	response.ResponseStatusAccept(c, roomnum)
}
