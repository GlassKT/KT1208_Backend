package models

import (
	"GlassKT/database"
)

func StoreImg(id, imageName string) error {
	User := &database.User{}
	imguser := database.User{ID: id, IMGNAME: imageName}
	err := database.DB.Model(&User).Updates(imguser).Error
	if err != nil {
		return err
	}
	return nil
}
