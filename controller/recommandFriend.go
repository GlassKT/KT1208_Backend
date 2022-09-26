package controller

import (
	"fmt"
	"web_test/database"
	"web_test/models"
	"web_test/modules"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

// 친구의 친구를 추천하는 경우
func RecommandFriendFriends(c *gin.Context) {

	id := c.Query("id")

	user := []models.WebFriendHobbyInfo{}

	err := database.DB.Model(&models.Friend{}).Omit("hobby").Where("user = ?", id).Find(&user).Error
	if err != nil {
		response.ResponseBadRequest(c, "recommand friend friends find erorr")
		return
	}

	friends := []string{}

	fmt.Println(user)

	for _, v := range user {
		friends = append(friends, v.Friend)
	}

	fmt.Println(friends)

	hobbies := modules.GetHobbyStringSlice(friends)

	for i, v := range user {
		for _, v2 := range hobbies {
			if v2.User == v.Friend {
				user[i].Hobby = append(user[i].Hobby, v2.Hobby)
			}
		}
	}

	response.ResponseList(c, user)
}

// 그냥 랜덤으로 추천
func RecommandFriend(c *gin.Context) {

	user := []models.WebUserInfo{}

	_ = database.DB.Raw("SELECT id FROM users ORDER BY rand() LIMIT 20").Scan(&user) // 10명까지 랜덤으로 추천

	friends := []string{}

	for _, v := range user {
		friends = append(friends, v.Id)
	}

	fmt.Println(user)
	fmt.Println(friends)

	hobbies := modules.GetHobbyStringSlice(friends)

	for i, v := range user {
		for _, v2 := range hobbies {
			if v2.User == v.Id {
				user[i].Hobby = append(user[i].Hobby, v2.Hobby)
			}
		}
	}

	response.ResponseList(c, user)
}
