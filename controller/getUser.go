package controller

import (
	"web_test/database"
	"web_test/models"
	"web_test/modules"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	id := c.Query("id")

	user := &models.GetUserInfo{}

	if err := database.DB.Model(&models.User{}).Omit("hobby", "Hobby").Where("id = ?", id).Find(user).Error; err != nil {
		response.ResponseBadRequest(c, "get user get User error")
		return
	}

	hobby := modules.GetHobby(user.Id)

	for _, v := range hobby {
		user.Hobby = append(user.Hobby, v.Hobby)
	}

	response.ResponseList(c, *user)
}
