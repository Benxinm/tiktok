package cache

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
	"time"
)

const (
	CommentCountKey = "comment_count"
)

func GetCommentCountKey(vid int64) string {
	return fmt.Sprintf("%s_%d", CommentCountKey, vid)
}

func SetCommentCount(ctx context.Context, vid int64, count int64) error {
	err := RedisClient.Set(ctx, GetCommentCountKey(vid), ctx, time.Hour).Err()
	if err != nil {
		klog.Error(err)
		return err
	}
	return err
}

func GetCommentCount(ctx context.Context, vid int64) (int64, error) {
	result, err := RedisClient.Get(ctx, GetCommentCountKey(vid)).Result()
	if err != nil {
		return -1, err
	}
	count, _ := strconv.ParseInt(result, 10, 64)
	return count, nil
}

func AddCommentCount(ctx context.Context, vid int64) error {
	err := RedisClient.Incr(ctx, GetCommentCountKey(vid)).Err()
	if err != nil {
		klog.Error(err)
	}
	return err
}
