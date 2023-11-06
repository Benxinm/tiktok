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

func Follow(ctx context.Context, uid int64, tuid int64) error {
	followKey := fmt.Sprintf("%s_%v", FollowKey, uid)
	followerKey := fmt.Sprintf("%s_%v", FollowerKey, tuid)
	follow, err := IsFollow(ctx, uid, tuid)
	if follow {
		return err
	}
	err = RedisClient.SAdd(ctx, followKey, strconv.FormatInt(tuid, 10)).Err()
	if err != nil {
		klog.Infof("err: %v", err)
		return err
	}
	err = RedisClient.SAdd(ctx, followerKey, strconv.FormatInt(uid, 10)).Err()
	if err != nil {
		klog.Infof("err: %v", err)
		return err
	}
	return nil
}

func UnFollow(ctx context.Context, uid int64, tuid int64) error {
	followKey := fmt.Sprintf("%s_%v", FollowKey, uid)
	followerKey := fmt.Sprintf("%s_%v", FollowerKey, tuid)
	follow, err := IsFollow(ctx, uid, tuid)
	if !follow {
		return err
	}
	err = RedisClient.SRem(ctx, followKey, strconv.FormatInt(tuid, 10)).Err()
	if err != nil {
		klog.Infof("err: %v", err)
		return err
	}
	err = RedisClient.SRem(ctx, followerKey, strconv.FormatInt(uid, 10)).Err()
	if err != nil {
		klog.Infof("err: %v", err)
		return err
	}
	return nil
}

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

func IsFollow(ctx context.Context, uid int64, tuid int64) (bool, error) {
	key := fmt.Sprintf("%s_%v", FollowKey, uid)
	isFollow, err := RedisClient.SIsMember(ctx, key, strconv.FormatInt(tuid, 10)).Result()
	if err != nil {
		klog.Infof("err: %v", err)
		return isFollow, err
	}
	return isFollow, nil
}

func FriendList(ctx context.Context, uid int64) ([]int64, error) {
	friendList := make([]int64, 20)
	list, err := FollowList(ctx, uid)
	if err != nil {
		klog.Info("err: %v", err)
		return nil, err
	}
	key := fmt.Sprintf("%s_%v", FollowerKey, uid)
	for _, followId := range list {
		result, err := RedisClient.SIsMember(ctx, key, followId).Result()
		if err != nil {
			klog.Info("err: %v", err)
			return nil, err
		} else if !result {
			continue
		}
		friendList = append(friendList, followId)
	}
	return friendList, err
}
