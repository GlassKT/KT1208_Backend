package models

import (
	db "GlassKT/database"
)

func MakeUser(id, pw, name, email string) error {
	User := &db.User{ID: id, PW: pw, NAME: name, EMAIL: email}
	err := db.DB.Omit("createAt", "birth").Create(User).Error
	return err
}
