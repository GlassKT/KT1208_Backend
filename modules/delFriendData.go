package modules

import (
	"web_test/database"
	"web_test/models"
)

func DelFriendData(f *models.WebFriendInfo, model interface{}) error {

	var err error

	if err = database.DB.Where("user = ? and friend = ?", f.User, f.Friend).Delete(model).Error; err != nil {
		return err
	}

	if err = database.DB.Where("user = ? and friend = ?", f.Friend, f.User).Delete(model).Error; err != nil {
		return err
	}

	return nil
}
