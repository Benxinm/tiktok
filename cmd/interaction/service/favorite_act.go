package service

import (
	"github.com/benxinm/tiktok/cmd/interaction/dal/cache"
	"github.com/benxinm/tiktok/cmd/interaction/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/interaction"
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
	_, err = cache.GetVideoLikeCount(s.ctx, req.VideoId)
	if err != nil {
		err := cache.SetVideoLikeCount(s.ctx, req.VideoId, 0)
		if err != nil {
			return err
		}
	}
	_, err = cache.GetVideoLikeCount(s.ctx, req.VideoId)
	if err != nil {
		err := cache.SetVideoLikeCount(s.ctx, req.VideoId, 0)
		if err != nil {
			return err
		}
	}
	err = cache.AddVideFavoriteCount(s.ctx, req.VideoId, uid)
	return err
}
