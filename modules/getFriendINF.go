package modules

import (
	"fmt"
	"web_test/database"
	"web_test/models"
)

func GetFriendINFSlice(name []string) []models.GetFriendInfo {

	friend := []models.GetFriendInfo{}

	for _, v := range name {

		getFriendTemp := &models.GetFriendInfo{}

		err := database.DB.Omit("hobby", "Hobby").Model(&models.User{}).Where("id = ?", v).Find(getFriendTemp).Error
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			friend = append(friend, *getFriendTemp)
		}

	}
	return friend
}
