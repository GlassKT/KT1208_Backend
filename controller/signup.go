package controller

import (
	_ "database/sql"
	"web_test/database"
	"web_test/helper"
	"web_test/models"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {

	signupInfo := &models.User{} // signupInfo로 create를 하고 싶었으나 에러 발생으로 실패 updates는 가능한데 안되는것이 의문
	// signupInfo := &models.LoginInfo{} // signupInfo로 create를 하고 싶었으나 에러 발생으로 실패 updates는 가능한데 안되는것이 의문

	err := c.Bind(signupInfo) // bind
	if err != nil {
		response.ResponseBadRequest(c, "signup bind error")
		return
	}

	if err := database.DB.Model(&models.User{}).Where("id = ?", signupInfo.Id).Find(signupInfo).Error; err != nil { // find id
		response.ResponseBadRequest(c, "signup where error")
		return
	}

	if signupInfo.Pw, err = helper.HashPassword(signupInfo.Pw); err != nil { // hash pwd
		response.ResponseBadRequest(c, "signup hashpwd error")
		return
	}

	if err = database.DB.Select("id", "pw", "name", "email").Create(signupInfo).Error; err != nil { // create user
		response.ResponseBadRequest(c, "signup create error")
		return
	}

	// if err = database.DB.Model(&models.User{}).Select("id", "pw", "name", "email").Create(signupInfo).Error; err != nil { // create user
	// 	response.ResponseBadRequest(c, "signup create error")
	// 	return
	// }

	response.ResponseStatusAccept(c, "signup success")
}
