package service

import (
	"github.com/benxinm/tiktok/cmd/interaction/dal/cache"
	"github.com/benxinm/tiktok/cmd/interaction/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) GetUserFavoriteCount(req *interaction.UserFavoriteCountRequest) (int64, error) {
	count, err := cache.GetUserFavoriteCount(s.ctx, req.UserId)
	if err != nil {
		return 0, err
	}
	if count == 0 {
		count, err = db.GetUserFavouriteCount(s.ctx, req.UserId)
	}
	return count, err
}
