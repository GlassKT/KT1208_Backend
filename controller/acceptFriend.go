package controller

import (
	"web_test/database"
	"web_test/models"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func AcceptFriend(c *gin.Context) {

	webFriendInfo := &models.WebFriendInfo{}

	err := c.Bind(webFriendInfo)
	if err != nil {
		response.ResponseBadRequest(c, "accept friend bind error")
		return
	}

	if webFriendInfo.User == "" || webFriendInfo.Friend == "" {
		response.ResponseBadRequest(c, "유저나 친구의 아이디가 제대로 오지 않음")
		return
	}

	if err = database.DB.Model(&models.Friend{}).Where("user = ? and friend = ?", webFriendInfo.User, webFriendInfo.Friend).Update("status", "allow").Error; err != nil { // 첫번째 친구 수락
		response.ResponseBadRequest(c, "accept friend first update error")
		return
	}

	if err = database.DB.Model(&models.Friend{}).Where("user = ? and friend = ?", webFriendInfo.Friend, webFriendInfo.User).Update("status", "allow").Error; err != nil { // 두번째 친구 수락
		response.ResponseBadRequest(c, "accept friend first update error")
		return
	}

	response.ResponseStatusAccept(c, "accept friend success")
}
