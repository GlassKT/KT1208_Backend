package modules

import (
	"context"
	"time"
	"web_test/models"
	"web_test/redis"
)

// 토큰 정보를 redis에 저장
func StoreAuth(td *models.TokenDetails) error {

	at := time.Unix(td.AtExpires, 0)
	now := time.Now()

	errAccess := redis.Client.Set(context.Background(), td.AccessUuid, td.AccessToken, at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	return nil
}
