package modules

import (
	"time"
	"web_test/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var accessString = []byte("secret") // access token의 key

// 토큰 제작
func CreateJWT(id string) (*models.TokenDetails, error) {

	var err error

	td := &models.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Hour * 1).Unix() // 엑세스 토큰 15분
	td.AccessUuid = uuid.NewString()

	// 엑세스 토큰
	atClaims := jwt.MapClaims{}
	atClaims["id"] = id
	atClaims["exp"] = td.AtExpires
	atClaims["access_uuid"] = td.AccessUuid

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims) // claims를 담아 토큰 만들기
	td.AccessToken, err = at.SignedString([]byte(accessString))
	if err != nil {
		return nil, err
	}

	return td, nil
}
