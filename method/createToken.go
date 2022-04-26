package method

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(id string) (string, error) {

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	access := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, err := access.SignedString([]byte())

}
