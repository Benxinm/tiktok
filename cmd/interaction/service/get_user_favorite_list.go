package service

import (
	"errors"
	"github.com/benxinm/tiktok/cmd/interaction/dal/cache"
	"github.com/benxinm/tiktok/cmd/interaction/dal/db"
	"github.com/benxinm/tiktok/cmd/interaction/rpc"
	"github.com/benxinm/tiktok/kitex_gen/interaction"
	"github.com/benxinm/tiktok/kitex_gen/video"
	"gorm.io/gorm"
)

func (s *InteractionService) FavoriteList(req *interaction.FavoriteListRequest) ([]*video.Video, error) {
	vids, err := cache.GetUserFavoriteVideos(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if len(vids) != 0 {
		return rpc.GetFavoriteVideoList(s.ctx, &video.GetFavoriteVideoInfoRequest{
			VideoId: vids,
			Token:   req.Token,
		})
	}
	vids, err = db.GetFavouriteVideosByUid(s.ctx, req.UserId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	err = cache.UpdateFavoriteVideoList(s.ctx, req.UserId, vids)
	if err != nil {
		return nil, err
	}
	if len(vids) == 0 {
		return nil, nil
	}
	return rpc.GetFavoriteVideoList(s.ctx, &video.GetFavoriteVideoInfoRequest{
		VideoId: vids,
		Token:   req.Token,
	})
}
