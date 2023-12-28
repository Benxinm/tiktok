package service

import (
	"context"
	"github.com/benxinm/tiktok/cmd/interaction/dal/cache"
	"github.com/benxinm/tiktok/cmd/interaction/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/interaction"
	"github.com/cloudwego/kitex/pkg/klog"
	"math/rand"
	"time"
)

func (s *InteractionService) Favorite(req *interaction.FavoriteActionRequest, uid int64) error {
	exist, err := db.IsFavoriteExist(s.ctx, uid, req.VideoId)
	if !exist && err == nil {
		err = db.CreateFavorite(s.ctx, &db.Favorite{
			VideoID: req.VideoId,
			UserID:  uid,
			Status:  1,
		})
		if err != nil {
			return err
		}
	} else if exist {
		err = db.UpdateFavoriteStatus(s.ctx, uid, req.VideoId, 1)
	}
	// 既然已经到喜欢和不喜欢这一步了，cache中必定有video count的缓存 所以不需要再进一步检验key是否存在
	err = cache.AddVideFavoriteCount(s.ctx, req.VideoId, uid)
	go delayWrite(s.ctx, req.VideoId)
	return err
}

func delayWrite(ctx context.Context, vid int64) {
	flag, err := cache.IsVideoInSet(ctx, vid)
	if flag {
		return
	}
	duration := time.Duration(3000 + rand.Intn(1000))
	if !flag || err != nil {
		err := cache.AddVideoSet(ctx, vid, duration)
		if err != nil {
			klog.Warnf("%v add set failed", vid)
		}
	}
	time.Sleep(duration)
	count, err := cache.GetVideoLikeCount(ctx, vid)
	if err != nil {
		return
	}
	err = db.UpdateVideoFavouriteCount(ctx, vid, count)
	cache.RemoveVideoSet(ctx, vid)
}
