package modules

import (
	"context"
	"web_test/redis"
)

func DeleteAuth(dtk string) (int64, error) {

	dresult, err := redis.Client.Del(context.Background(), dtk).Result()
	if err != nil {
		return 0, err
	}
	return dresult, nil
}
