package modules

import (
	"example.com/m/database"
	"example.com/m/models"
)

func StoreImg(id, imageName string) error {
	User := &models.User{}
	imguser := models.User{ID: id, IMGNAME: imageName}
	err := database.DB.Model(&User).Updates(imguser).Error
	if err != nil {
		return err
	}
	return nil
}
