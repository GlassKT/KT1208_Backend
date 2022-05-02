package models

import (
	db "GlassKT/database"
)

func MakeUser(id, pw, name, email string) error {
	User := &db.User{ID: id, PW: pw, NAME: name, EMAIL: email}
	err := db.DB.Create(User).Error
	return err
}
