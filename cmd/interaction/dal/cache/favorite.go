package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

const (
	UserLikeKey       = "user_like"
	VideoLikeCountKey = "video_like_count"
)

func UserFavoriteKey(uid int64) string {
	return fmt.Sprintf("%s_%d", UserLikeKey, uid)
}
func VideoFavoriteCountKey(vid int64) string {
	return fmt.Sprintf("%s_%d", VideoLikeCountKey, vid)
}

func IsVideoFavoriteExist(ctx context.Context, vid int64, uid int64) (bool, error) {
	result, err := RedisClient.SIsMember(ctx, UserFavoriteKey(uid), strconv.FormatInt(vid, 10)).Result()
	if err != nil {
		return false, err
	}
	return result, err
}

func AddVideFavoriteCount(ctx context.Context, vid int64, uid int64) error {
	pipe := RedisClient.TxPipeline()

	if err := pipe.SAdd(ctx, UserFavoriteKey(uid), strconv.FormatInt(vid, 10)).Err(); err != nil {
		return err
	}

	if err := pipe.Incr(ctx, VideoFavoriteCountKey(vid)).Err(); err != nil {
		return err
	}
	if err := pipe.Expire(ctx, VideoFavoriteCountKey(vid), time.Hour*1).Err(); err != nil {
		return err
	}
	_, err := pipe.Exec(ctx)
	return err
}

func ReduceVideoLikeCount(ctx context.Context, vid int64, uid int64) error {
	pipe := RedisClient.TxPipeline()

	if err := pipe.SRem(ctx, UserFavoriteKey(uid), strconv.FormatInt(vid, 10)).Err(); err != nil {
		return err
	}

	if err := pipe.Decr(ctx, VideoFavoriteCountKey(vid)).Err(); err != nil {
		return err
	}
	if err := pipe.Expire(ctx, VideoFavoriteCountKey(vid), time.Hour*1).Err(); err != nil {
		return err
	}
	_, err := pipe.Exec(ctx)
	return err
}

func GetVideoLikeCount(ctx context.Context, vid int64) (bool, int64, error) {
	count, err := RedisClient.Get(ctx, VideoFavoriteCountKey(vid)).Result()
	if err == redis.Nil {
		return false, 0, err
	}
	if err != nil {
		return true, 0, err
	}
	likes, _ := strconv.ParseInt(count, 10, 64)
	return true, likes, nil
}

func GetUserFavoriteVideos(ctx context.Context, uid int64) ([]int64, error) {
	items, err := RedisClient.SMembers(ctx, UserFavoriteKey(uid)).Result()
	if err != nil {
		return nil, err
	}
	vids := make([]int64, len(items))
	for _, item := range items {
		vid, _ := strconv.ParseInt(item, 10, 64)
		vids = append(vids, vid)
	}
	return vids, nil
}

func UpdateFavoriteVideoList(ctx context.Context, uid int64, vids []int64) error {
	var err error
	for i := 0; i < len(vids); i++ {
		err = RedisClient.SAdd(ctx, UserFavoriteKey(uid), strconv.FormatInt(vids[i], 10)).Err()
	}
	return err
}

func GetUserFavoriteCount(ctx context.Context, uid int64) (int64, error) {
	count, err := RedisClient.SCard(ctx, UserFavoriteKey(uid)).Result()
	if err != nil {
		return 0, err
	}
	return count, nil
}