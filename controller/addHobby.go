package controller

import (
	"fmt"
	"web_test/database"
	"web_test/models"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func AddHobby(c *gin.Context) {

	webHobbyInfo := &models.WebHobbyInfo{}

	sliceHobby := []models.Hobby{}

	err := c.Bind(webHobbyInfo)
	if err != nil {
		response.ResponseBadRequest(c, "add hobby bind error")
		fmt.Println(err)
		return
	}

	for _, v := range webHobbyInfo.Hobby {
		a := models.Hobby{User: webHobbyInfo.User, Hobby: v}
		sliceHobby = append(sliceHobby, a)
	}

	if err = database.DB.Create(&sliceHobby).Error; err != nil { // 배열 슬라이스에 담아 한번에 query
		response.ResponseBadRequest(c, "add hobby create error")
		return
	}

	response.ResponseStatusAccept(c, "add hobby success")

	// 처음 코드, 문제점 : 취미가 늘어남에 따라 query가 지나치게 많아질 수 있음

	// hobbyInfo := &models.WebHobbyInfo{}

	// hobby := &models.Hobby{}

	// err := c.Bind(hobbyInfo)
	// if err != nil {
	// 	response.ResponseBadRequest(c, "add hobby bind error")
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(hobbyInfo)

	// hobby.User = hobbyInfo.User

	// for _, v := range hobbyInfo.Hobby {
	// 	hobby.Hobby = v
	// 	if err = database.DB.Create(hobby).Error; err != nil {
	// 		response.ResponseBadRequest(c, "add hobby create error")
	// 		return
	// 	}
	// }

	// response.ResponseStatusAccept(c, "add hobby success")
}
