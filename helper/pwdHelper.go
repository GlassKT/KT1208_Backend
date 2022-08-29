package helper

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pw string) string {

	bytes, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(bytes)
}

func CheckPasswordHash(hashVal, userPw string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashVal), []byte(userPw))

	fmt.Println(err)

	if err != nil {
		return false
	} else {
		return true
	}
}
