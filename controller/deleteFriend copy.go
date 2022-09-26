package controller

import (
	"web_test/models"
	"web_test/modules"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func DeleteBlockFriend(c *gin.Context) {

	webFriendInfo := &models.WebFriendInfo{}

	err := c.Bind(webFriendInfo)
	if err != nil {
		response.ResponseBadRequest(c, "delete block friend bind error")
		return
	}

	if err = modules.DelFriendData(webFriendInfo, &models.BlockFriend{}); err != nil {
		response.ResponseBadRequest(c, "delete block friend del friend data error")
		return
	}

	response.ResponseStatusAccept(c, "delete friend success")
}
