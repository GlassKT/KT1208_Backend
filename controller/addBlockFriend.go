package controller

import (
	"web_test/database"
	"web_test/models"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func AddBlockFriend(c *gin.Context) {

	blocFriend := &models.BlockFriend{}

	err := c.Bind(blocFriend)
	if err != nil {
		response.ResponseBadRequest(c, "block friend bind error")
		return
	}

	if err = database.DB.Where("user = ? and friend = ?", blocFriend.User, blocFriend.Friend).Find(&models.Friend{}).Error; err != nil { // 친구 여부
		response.ResponseBadRequest(c, "친구 상태 입니다")
		return
	}

	if err = database.DB.Where("user = ? and friend = ?", blocFriend.User, blocFriend.Friend).Find(&models.BlockFriend{}).Error; err != nil { // 친구 차단 여부
		response.ResponseBadRequest(c, "친구 차단이 이미 되어 있습니다")
		return
	}

	if err = database.DB.Create(blocFriend).Error; err != nil {
		response.ResponseBadRequest(c, "block friend first create error")
		return
	}

	temp := blocFriend.User
	blocFriend.User = blocFriend.Friend
	blocFriend.Friend = temp

	if err = database.DB.Create(blocFriend).Error; err != nil {
		response.ResponseBadRequest(c, "block friend second create error")
		return
	}

	response.ResponseStatusAccept(c, "block friend success")
}
