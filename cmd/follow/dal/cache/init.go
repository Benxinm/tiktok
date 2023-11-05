package cache

import (
	"context"
	"github.com/benxinm/tiktok/pkg/constants"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       constants.FollowRedis,
	})
	_, err := RedisClient.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}

}
