package modules

import (
	"web_test/database"
	"web_test/models"
)

/*func GetHobby(friend []models.GetFriendInfo) []models.Hobby {

	hobby := []models.Hobby{}

	for _, v := range friend {

		hb := &models.Hobby{}

		err := database.DB.Model(&models.Hobby{}).Where("user = ?", v.Id).Find(hobby).Error
		if err != nil {
			continue
		}

		hobby = append(hobby, *hb)
	}

	return hobby
}*/

func GetHobbySlice(friend []models.GetFriendInfo) []models.Hobby {

	hobby := []models.Hobby{}

	for _, v := range friend {

		hb := []models.Hobby{}

		err := database.DB.Model(&models.Hobby{}).Where("user = ?", v.Id).Find(&hb).Error
		if err != nil {
			continue
		}

		for _, v := range hb {
			hobby = append(hobby, v)
		}

	}

	return hobby
}

func GetHobby(user string) []models.Hobby {

	hobby := []models.Hobby{}

	err := database.DB.Model(&models.Hobby{}).Where("user = ?", user).Find(&hobby).Error
	if err != nil {
		return nil
	}

	return hobby
}

func GetHobbyStringSlice(friend []string) []models.Hobby {

	hobby := []models.Hobby{}

	for _, v := range friend {

		hb := []models.Hobby{}

		err := database.DB.Model(&models.Hobby{}).Where("user = ?", v).Find(&hb).Error
		if err != nil {
			continue
		}

		for _, v := range hb {
			hobby = append(hobby, v)
		}
	}

	return hobby
}
