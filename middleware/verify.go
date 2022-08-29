package middleware

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ResponseBad(c *gin.Context, code int, errM interface{}) {
	c.JSON(code, gin.H{
		"message": errM,
	})
}

func TokenAuthMiddkeware() gin.HandlerFunc {

	return func(c *gin.Context) {

		token, err := c.Request.Cookie("Authorization")
		if err != nil {
			ResponseBad(c, 401, "token had not founded")
			c.Abort()
			return
		}

		tokenValue := token.Value

		claims := jwt.MapClaims{}

		token1, err := jwt.ParseWithClaims(tokenValue, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte([]byte("secret")), nil
		})

		log.Println("parsed token = ", token1)

		c.Next()
	}

}
