package controller

import (
	"log"
	"web_test/database"
	"web_test/models"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

type resMyLastChatting struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	ImgName string `json:"image"`
	NumRoom string `json:"roomnum"`

	LastChatting string `json:"message"`
}

type getRoomnum struct {
	RoomNum string
}

type messageUser struct {
	UserId  string
	Content string
	NumRoom string
}

type user struct {
	Id      string
	Name    string
	ImgName string
}

func GetMyChatting(c *gin.Context) {

	userId := c.Query("id")

	roomNumbers := []getRoomnum{}
	lastChattings := []resMyLastChatting{}

	err := database.DB.Model(&models.ChattingRoom{}).Where("user_id", userId).Find(&roomNumbers).Error
	if err != nil {
		log.Println(err)
		response.ResponseBadRequest(c, "get my chatting first find error")
		return
	}

	for _, v := range roomNumbers {

		msgUser := &messageUser{}
		userUser := &user{}

		err = database.DB.Model(&models.Message{}).Where("num_room", v.RoomNum).Order("create_at DESC").Find(msgUser).Error
		if err != nil {
			log.Println(err)
			response.ResponseBadRequest(c, "get my chatting second find error")
			return
		}

		err = database.DB.Model(&models.User{}).Where("user_id = ?", msgUser.UserId).Find(userUser).Error
		if err != nil {
			log.Println(err)
			response.ResponseBadRequest(c, "get my chatting second find error")
			return
		}

		lastChat := &resMyLastChatting{
			Id:           msgUser.UserId,
			LastChatting: msgUser.Content,
			NumRoom:      msgUser.NumRoom,
			Name:         userUser.Name,
			ImgName:      userUser.ImgName,
		}

		lastChattings = append(lastChattings, *lastChat)
	}

	response.ResponseList(c, lastChattings)
}
