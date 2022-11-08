package controller

import (
	"web_test/database"
	"web_test/models"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

type chattingHistory struct {
	NumRoom  string `json:"roomnum"`
	UserId   string `json:"id"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	Createat string `json:"date"`
}

func GetChattingHistory(c *gin.Context) {

	roomnum := c.Query("roomnum")

	chatHis := []chattingHistory{}

	err := database.DB.Model(&models.Message{}).Where("num_room = ?", roomnum).Order("messagescol DESC").Find(&chatHis).Error
	if err != nil {
		response.ResponseBadRequest(c, "get chatting history find error")
		return
	}

	response.ResponseList(c, chatHis)
}
