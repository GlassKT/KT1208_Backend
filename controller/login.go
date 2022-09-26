package controller

import (
	"web_test/database"
	"web_test/helper"
	"web_test/models"
	"web_test/modules"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	loginInfo := &models.LoginInfo{}

	err := c.Bind(loginInfo) //bind
	if err != nil {
		response.ResponseBadRequest(c, "login bind error")
		return
	}

	orgPwd := loginInfo.Pw

	if err = database.DB.Model(&models.User{}).Where("id = ?", loginInfo.Id).Find(loginInfo).Error; err != nil { // id check
		response.ResponseBadRequest(c, "login id is not true")
		return
	}

	if !helper.CheckPasswordHash(loginInfo.Pw, orgPwd) { // pw check
		response.ResponseBadRequest(c, "login pw is not true")
		return
	}

	tk, err := modules.CreateJWT(loginInfo.Id)

	if err != nil { // create token
		response.ResponseBadRequest(c, "login create token error")
		return
	}

	if err := modules.StoreAuth(tk); err != nil { // store token
		response.ResponseBadRequest(c, "login auth a error")
		return
	}

	response.ResponseToken(c, tk.AccessToken, loginInfo.Id)
}
