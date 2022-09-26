package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pw string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(hashVal, userPw string) bool { // hash값과 일반 string을 받아 비교

	err := bcrypt.CompareHashAndPassword([]byte(hashVal), []byte(userPw))

	if err != nil {
		return false
	} else {
		return true
	}
}
