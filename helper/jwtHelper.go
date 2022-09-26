package helper

import (
	"fmt"
	"web_test/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// token 추출
func ExtractToken(c *gin.Context) string {

	token := c.Request.Header.Get("Authorization")

	if token == "" {
		return ""
	}

	return token
}

// token의 파싱 방법을 찾아 비교하고 token형태로 반환
func VerifyToken(c *gin.Context) (*jwt.Token, error) {

	tokenString := ExtractToken(c)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singing method : %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// token 유효성 검사
func TokenVaild(c *gin.Context) error {

	token, err := VerifyToken(c)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// token에서 redis에 접근하기위 한 access uuid와 유저의 id를 추출
func ExtractTokenMetadata(c *gin.Context) (*models.AccessDetails, error) {

	token, err := VerifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId := fmt.Sprintf("%s", claims["id"])

		if err != nil {
			return nil, err
		}
		return &models.AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}
