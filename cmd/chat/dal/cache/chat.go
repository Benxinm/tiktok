package cache

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

func CreateMessage(ctx context.Context, key string, revkey string, stamp_key int64, field string) error {
	// 先判断是否key是否存在，如果存在则判断过期时间是否小于十天，小于则加时间，大于则不加时间
	if ok := MessageExist(ctx, key); ok != 0 {
		err := RedisClient.HSet(ctx, key, stamp_key, field).Err()
		if err != nil {
			klog.Error(err)
			return err
		}
		lastTime, err := RedisClient.TTL(ctx, key).Result()
		if err != nil {
			klog.Error(err)
			return err
		}
		if lastTime.Seconds() < 60*60*24*10 {
			err = RedisClient.Expire(ctx, key, time.Hour*24*30).Err()
			if err != nil {
				klog.Error(err)
				return err
			}
			err = RedisClient.Expire(ctx, revkey, time.Hour*24*30).Err()
			if err != nil {
				klog.Error(err)
				return err
			}
			return nil
		}
	}
	// 不存在则直接add然后加时间
	err := RedisClient.HSet(ctx, key, stamp_key, field).Err()
	if err != nil {
		klog.Error(err)
		return err
	}
	err = RedisClient.Expire(ctx, key, time.Hour*24*30).Err()
	if err != nil {
		klog.Error(err)
		return err
	}
	err = RedisClient.Expire(ctx, revkey, time.Hour*24*30).Err()
	if err != nil {
		klog.Error(err)
		return err
	}
	return nil
}

func GetMessage(ctx context.Context, key string) ([]string, error) {
	mem, err := RedisClient.HVals(ctx, key).Result()
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	return mem, nil
}

func MessageExist(ctx context.Context, key string) int64 {
	ok, err := RedisClient.Exists(ctx, key).Result()
	if err != nil {
		klog.Error(err)
		return 0
	}
	return ok
}
