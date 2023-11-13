package service

import (
	"github.com/benxinm/tiktok/cmd/interaction/rpc"
	"github.com/benxinm/tiktok/kitex_gen/interaction"
	"github.com/benxinm/tiktok/kitex_gen/video"
	"golang.org/x/sync/errgroup"
)

func (s *InteractionService) GetUserTotalFavorited(req *interaction.UserTotalFavoritedRequest) (int64, error) {
	vids, err := rpc.GetUserVideoList(s.ctx, &video.GetVideoIDByUidRequest{
		UserId: req.UserId,
		Token:  req.Token,
	})
	if err != nil {
		return 0, err
	}
	var total int64 = 0
	var eg errgroup.Group
	for _, vid := range vids {
		eg.Go(func() error {
			count, err := s.GetVideoFavoriteCount(&interaction.VideoFavoritedCountRequest{
				VideoId: vid,
				Token:   req.Token,
			})
			total += count
			return err
		})
	}
	if err = eg.Wait(); err != nil {
		return total, err
	}
	return total, nil
}
