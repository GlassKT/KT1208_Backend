package method

import (
	db "GlassKT/database"
	"time"
)

func MakeUser(id, pw, name, email string, crt time.Time) error {
	User := &db.User{ID: id, PW: pw, NAME: name, EMAIL: email, CREATE_AT: crt}
	err := db.DB.Create(User).Error
	return err
}
