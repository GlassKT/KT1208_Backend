package controller

import (
	"web_test/helper"
	"web_test/modules"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func LogOut(c *gin.Context) {

	tk, err := helper.ExtractTokenMetadata(c) // token 정보 추출
	if err != nil {
		response.ResponseBadRequest(c, "logout toekn is not found")
		return
	}

	dresult, err := modules.DeleteAuth(tk.AccessUuid) // redis token 삭제
	if err != nil || dresult == 0 {
		response.ResponseBadRequest(c, "logout delete auth error")
		return
	}

	response.ResponseStatusAccept(c, "logout success")
}
