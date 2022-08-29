package modules

import (
	db "example.com/m/database"
	"example.com/m/models"
)

func MakeUser(id, pw, name, email string) error {
	User := &models.User{ID: id, PW: pw, NAME: name, EMAIL: email}
	err := db.DB.Omit("createAt", "birth").Create(User).Error
	return err
}
