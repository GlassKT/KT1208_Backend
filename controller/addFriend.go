package controller

import (
	"web_test/database"
	"web_test/models"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func AddFriend(c *gin.Context) {

	friend := &models.Friend{}

	err := c.BindJSON(friend) // friend creatr
	if err != nil {
		response.ResponseBadRequest(c, "add friend bind error")
		return
	}

	if err = database.DB.Where("user = ? and friend = ?", friend.User, friend.Friend).First(&models.Friend{}).Error; err == nil { // 친구 존재 확인
		response.ResponseBadRequest(c, "친구 추가가 이미 되어 있습니다")
		return
	}

	friend.Status = "wait" // 친구 신청을 받기 전이므로 wait

	if err = database.DB.Create(friend).Error; err != nil { // friend create first
		response.ResponseBadRequest(c, "add friend first create error")
		return
	}

	temp := friend.User
	friend.User = friend.Friend
	friend.Friend = temp

	if err = database.DB.Create(friend).Error; err != nil { // friend create second
		response.ResponseBadRequest(c, "add friend second create error")
		return
	}

	response.ResponseStatusAccept(c, "add friend accept")
}
