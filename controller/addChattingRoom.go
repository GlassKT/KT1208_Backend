package controller

import (
	"math/rand"
	"strconv"
	"time"
	"web_test/database"
	"web_test/models"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

type chattingRoom struct {
	User    string `json:"id"`
	RoomNum string `json:"roomnum"`
	Friend  string `json:"friendid"`
}

func AddChattingRoom(c *gin.Context) {

	var randnum string

	var flag = false

	chatRoom := &chattingRoom{}
	temp := &models.ChattingRoom{}

	createChat := &models.ChattingRoom{}
	createChat2 := &models.ChattingRoom{}

	userIDtemp := []models.ChattingRoom{}
	friendIDtemp := []models.ChattingRoom{}

	err := c.Bind(chatRoom)
	if err != nil {
		response.ResponseBadRequest(c, "add chatting room bind error")
		return
	}

	if chatRoom.User == "" || chatRoom.Friend == "" {
		response.ResponseBadRequest(c, "게스트나 오너가 비어있음")
		return
	}

	if err := database.DB.Model(&models.ChattingRoom{}).Where("user = ?", chatRoom.User).Find(&userIDtemp).Error; err != nil {
		response.ResponseBadRequest(c, "add chatting room first find error")
		return
	}

	if err := database.DB.Model(&models.ChattingRoom{}).Where("user = ?", chatRoom.Friend).Find(&friendIDtemp).Error; err != nil {
		response.ResponseBadRequest(c, "add chatting room second find error")
		return
	}

	for _, v := range userIDtemp {
		for _, v2 := range friendIDtemp {
			if v.RoomNum == v2.RoomNum {
				flag = true
				randnum = v.RoomNum
				break
			}
		}
	}

	if flag {
		response.ResponseStatusAccept(c, randnum)
		return
	}

	for {
		randnum = strconv.Itoa(rand.Intn(int(time.Now().Unix())))

		res := database.DB.Model(&models.ChattingRoom{}).Where("room_num = ?", randnum).Find(temp)

		if res.RowsAffected == 0 {
			break
		} else {
			continue
		}
	}

	createChat.User = chatRoom.User
	createChat.RoomNum = randnum

	createChat2.User = chatRoom.Friend
	createChat2.RoomNum = randnum

	if err = database.DB.Model(&models.ChattingRoom{}).Create(createChat).Error; err != nil {
		response.ResponseBadRequest(c, "add chatting room first create error")
		return
	}

	if err = database.DB.Model(&models.ChattingRoom{}).Create(createChat2).Error; err != nil {
		response.ResponseBadRequest(c, "add chatting room second create error")
		return
	}

	response.ResponseStatusAccept(c, randnum)
}
