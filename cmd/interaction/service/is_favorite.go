package service

import (
	"errors"
	"github.com/benxinm/tiktok/cmd/interaction/dal/cache"
	"github.com/benxinm/tiktok/cmd/interaction/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/interaction"
	"gorm.io/gorm"
)

func (s *InteractionService) IsFavorite(req *interaction.IsFavoriteRequest) (bool, error) {
	exist, err := cache.IsVideoFavoriteExist(s.ctx, req.VideoId, req.UserId)
	if err != nil {
		return exist, err
	}
	if exist {
		return exist, nil
	}
	exist, err = db.IsFavoriteExist(s.ctx, req.UserId, req.VideoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
