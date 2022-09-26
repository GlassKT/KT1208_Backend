package middleware

import (
	"web_test/helper"
	"web_test/modules"
	"web_test/response"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc { // token확인 middleware

	return func(c *gin.Context) {

		if err := helper.TokenVaild(c); err != nil {
			response.ResponseUnauthorized(c, "vaild error")
		}

		ad, err := helper.ExtractTokenMetadata(c)
		if err != nil {
			response.ResponseUnauthorized(c, "extract token error")
		}

		if _, err := modules.CheckAuth(ad); err != nil {
			response.ResponseUnauthorized(c, "fetch auth error")
		}

		c.Next()
	}
}
