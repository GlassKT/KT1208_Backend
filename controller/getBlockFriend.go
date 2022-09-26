package controller

import (
	"web_test/database"
	"web_test/models"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func GetBlockFriend(c *gin.Context) {

	webFriendInfo := []models.WebFriendInfo{}

	user := c.Query("user")

	var err error

	if err = database.DB.Model(&models.BlockFriend{}).Where("user = ?", user).Find(&webFriendInfo).Error; err != nil {
		response.ResponseBadRequest(c, "get block friend find error")
		return
	}

	blockFriends := []string{}

	for _, v := range webFriendInfo {
		blockFriends = append(blockFriends, v.Friend)
	}

	response.ResponseList(c, blockFriends)

}
