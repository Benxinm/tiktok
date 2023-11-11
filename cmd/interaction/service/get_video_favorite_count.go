package service

import (
	"errors"
	"github.com/benxinm/tiktok/cmd/interaction/dal/cache"
	"github.com/benxinm/tiktok/cmd/interaction/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/interaction"
	"github.com/redis/go-redis/v9"
)

func (s *InteractionService) GetVideoFavoriteCount(req *interaction.VideoFavoritedCountRequest) (int64, error) {
	count, err := cache.GetVideoLikeCount(s.ctx, req.VideoId)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			count, err = db.GetVideoFavouriteCount(s.ctx, req.VideoId)
			if err != nil {
				return 0, err
			}
			cache.SetVideoLikeCount(s.ctx, req.VideoId, count)
		} else {
			return 0, err
		}
	}
	return count, err
}
