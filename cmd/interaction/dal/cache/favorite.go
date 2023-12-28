package cache

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

const (
	UserLikeKey       = "user_like"
	VideoLikeCountKey = "video_like_count"
	VideoSetKey       = "video_set"
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
	//User favourite list
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
	//User favourite list
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

/**
To implement delay write to db
*/

func AddVideoSet(ctx context.Context, vid int64, expire time.Duration) error {
	err := RedisClient.SAdd(ctx, VideoSetKey, strconv.FormatInt(vid, 10)).Err()
	RedisClient.Expire(ctx, VideoSetKey, expire)
	return err
}

func IsVideoInSet(ctx context.Context, vid int64) (bool, error) {
	return RedisClient.SIsMember(ctx, VideoSetKey, strconv.FormatInt(vid, 10)).Result()
}

func RemoveVideoSet(ctx context.Context, vid int64) error {
	return RedisClient.SRem(ctx, VideoSetKey, strconv.FormatInt(vid, 10)).Err()
}

func SetVideoLikeCount(ctx context.Context, vid int64, count int64) error {
	return RedisClient.Set(ctx, VideoFavoriteCountKey(vid), count, time.Hour).Err()
}

func GetVideoLikeCount(ctx context.Context, vid int64) (int64, error) {
	count, err := RedisClient.Get(ctx, VideoFavoriteCountKey(vid)).Result()
	if err != nil {
		return 0, err
	}
	likes, _ := strconv.ParseInt(count, 10, 64)
	return likes, nil
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
