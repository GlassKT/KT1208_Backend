package modules

import (
	"context"
	"web_test/models"
	"web_test/redis"
)

// redis에 저장된 값이 없거나 토큰의 유효기간이 만료(자연적으로) 되었는지 확인
func CheckAuth(authD *models.AccessDetails) (string, error) {

	userId, err := redis.Client.Get(context.Background(), authD.AccessUuid).Result()
	if err != nil {
		return "", nil
	}

	return userId, nil
}
