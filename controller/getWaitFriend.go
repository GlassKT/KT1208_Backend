package controller

import (
	"web_test/database"
	"web_test/models"
	"web_test/modules"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func GetWaitFriend(c *gin.Context) {

	friend := []models.Friend{}

	friendList := []string{}

	user := c.Query("user")

	if err := database.DB.Where("user = ? and status = ?", user, "wait").Find(&friend).Error; err != nil { // 받은 id중 status가 wait인 것을 추출
		response.ResponseBadRequest(c, "get wait friend find error")
		return
	}

	for _, v := range friend { // friend의 친구 이름들을 friend list에 넣음
		friendList = append(friendList, v.Friend)
	}

	getFriendInfo := modules.GetFriendINFSlice(friendList) // friend의 이름들로 유저 정보들을 불러옴

	hobby := modules.GetHobbySlice(getFriendInfo)

	for i, v := range getFriendInfo {
		for _, v2 := range hobby {
			if v2.User == v.Id {
				getFriendInfo[i].Hobby = append(getFriendInfo[i].Hobby, v2.Hobby)
			}
		}
	}

	response.ResponseList(c, getFriendInfo)
}
