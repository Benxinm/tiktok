package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"strconv"
)

const (
	FollowKey   = "follow"
	FollowerKey = "follower"
)

func FollowList(ctx context.Context, uid int64) ([]int64, error) {
	followList := make([]int64, 15)
	key := fmt.Sprintf("%s_%v", FollowKey, uid)
	result, err := RedisClient.SMembers(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		klog.Errorf("")
	} else if err != nil {
		klog.Infof("follow list redis: %v", err)
	}
	for _, id := range result {
		id, _ := strconv.ParseInt(id, 10, 64)
		followList = append(followList, id)
	}
	return followList, nil
}

func FollowerList(ctx context.Context, uid int64) ([]int64, error) {
	followerList := make([]int64, 15)
	key := fmt.Sprintf("%s_%v", FollowerKey, uid)
	result, err := RedisClient.SMembers(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		klog.Errorf("")
	} else if err != nil {
		klog.Infof("follow list redis: %v", err)
	}
	for _, id := range result {
		id, _ := strconv.ParseInt(id, 10, 64)
		followerList = append(followerList, id)
	}
	return followerList, nil
}

func FollowCount(ctx context.Context, uid int64) (int64, error) {
	key := fmt.Sprintf("%s_%v", FollowKey, uid)
	result, err := RedisClient.SCard(ctx, key).Result()
	if err != nil {
		return -1, err
	}
	return result, nil
}

func FollowerCount(ctx context.Context, uid int64) (int64, error) {
	key := fmt.Sprintf("%s_%v", FollowerKey, uid)
	result, err := RedisClient.SCard(ctx, key).Result()
	if err != nil {
		return -1, nil
	}
	return result, nil
}
