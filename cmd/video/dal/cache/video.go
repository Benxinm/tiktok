package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/benxinm/tiktok/cmd/video/dal/db"
	"github.com/benxinm/tiktok/cmd/video/rpc"
	"github.com/benxinm/tiktok/kitex_gen/follow"
	"github.com/benxinm/tiktok/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
	"strings"
	"time"
)

func GetUserVideoFeed(ctx context.Context, uid int64, offset int64, token string) (vids []string, err error) {
	length, err := RedisClient.LLen(ctx, strconv.FormatInt(uid, 10)).Result()
	if err != nil {
		return nil, err
	}
	//if current length smaller than offset target then pull from regular feed or get sub video
	if length < offset+constants.OffsetTarget {
		lastComponent, err := RedisClient.LIndex(ctx, strconv.FormatInt(uid, 10), length).Result()
		if err != nil {
			return nil, err
		}
		//if current length doesn't fit request then use last update time to query follow's video
		if updateTime, ok := strings.CutSuffix("_", lastComponent); ok {
			uidsSub, _ := rpc.GetFollowList(ctx, &follow.FollowListRequest{
				UserId: uid,
				Token:  token,
			})
			for _, uidSub := range uidsSub {
				videos, _ := db.GetVideoByTimeUid(ctx, updateTime, uidSub)
				for _, vid := range videos {
					RedisClient.LPush(ctx, strconv.FormatInt(uid, 10), fmt.Sprintf("%v_%v", vid, time.Stamp))
				}
			}
		}
	}
	data, err := RedisClient.LRange(ctx, strconv.FormatInt(uid, 10), offset, offset+constants.OffsetTarget).Result()
	if err != nil || len(data) <= 0 {
		return nil, err
	}
	return data, nil
}

func AddVideoList(ctx context.Context, videoList []db.Video, latestTime int64) {
	videoJson, err := json.Marshal(videoList)
	if err != nil {
		klog.Error(err)
	}
	err = RedisClient.Set(ctx, strconv.FormatInt(latestTime, 10), videoJson, time.Minute*10).Err()
	if err != nil {
		klog.Error(err)
	}
}

func GetVideoList(ctx context.Context, latestTime int64) (videoList []db.Video, err error) {
	data, err := RedisClient.Get(ctx, strconv.FormatInt(latestTime, 10)).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(data), &videoList)
	if err != nil {
		return nil, err
	}
	return
}

func IsExistVideoInfo(ctx context.Context, latestTime int64) (exist int64, err error) {
	exist, err = RedisClient.Exists(ctx, strconv.FormatInt(latestTime, 10)).Result()
	return
}
