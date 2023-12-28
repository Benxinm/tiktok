package service

import (
	"errors"
	"github.com/benxinm/tiktok/cmd/interaction/dal/cache"
	"github.com/benxinm/tiktok/cmd/interaction/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/interaction"
	"gorm.io/gorm"
)

func (s *InteractionService) UnFavorite(req *interaction.FavoriteActionRequest, uid int64) error {
	exist, err := cache.IsVideoFavoriteExist(s.ctx, req.VideoId, uid)
	if err != nil {
		return err
	}
	if !exist {
		err := db.IsFavorited(s.ctx, uid, req.VideoId, 1)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	_, err = cache.GetVideoLikeCount(s.ctx, req.VideoId)
	if err != nil {
		return err
	}
	if err := cache.ReduceVideoLikeCount(s.ctx, req.VideoId, uid); err != nil {
		return err
	}
	go delayWrite(s.ctx, req.VideoId)
	return db.UpdateFavoriteStatus(s.ctx, uid, req.VideoId, 0)
}
