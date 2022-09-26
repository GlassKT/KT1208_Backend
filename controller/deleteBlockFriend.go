package controller

import (
	"web_test/models"
	"web_test/modules"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func DeleteFriend(c *gin.Context) {

	webFriendInfo := &models.WebFriendInfo{}

	err := c.Bind(webFriendInfo)
	if err != nil {
		response.ResponseBadRequest(c, "delete friend bind error")
		return
	}

	if err = modules.DelFriendData(webFriendInfo, &models.Friend{}); err != nil {
		response.ResponseBadRequest(c, "delete friend del friend data error")
		return
	}

	response.ResponseStatusAccept(c, "delete friend success")
}
