package controller

import (
	"fmt"
	"web_test/database"
	"web_test/models"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	id := c.Query("id")

	user := &models.User{}

	if err := database.DB.Model(&models.User{}).Where("id = ?", id).Preload("Hobbys").Find(user).Error; err != nil {
		fmt.Println(err)
		response.ResponseBadRequest(c, "get user get User error")
		return
	}

	//hobby := modules.GetHobby(user.Id)

	// for _, v := range hobby {
	// 	user.Hobbys = append(user.Hobbys, v)
	// }

	response.ResponseList(c, *user)
}
