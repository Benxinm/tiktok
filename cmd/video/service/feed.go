package service

import (
	"github.com/benxinm/tiktok/cmd/video/dal/cache"
	"github.com/benxinm/tiktok/kitex_gen/video"
)

func (s *VideoService) Feed(req *video.FeedRequest) {
	cache.GetUserVideoFeed(s.ctx, 0, 0, req.Token)
}
