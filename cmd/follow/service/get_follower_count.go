package service

import (
	"github.com/benxinm/tiktok/cmd/follow/dal/cache"
	"github.com/benxinm/tiktok/cmd/follow/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/follow"
)

func (s *FollowService) GetFollowerCount(req *follow.FollowerCountRequest) (int64, error) {
	count, err := cache.FollowerCount(s.ctx, req.UserId)
	if err != nil {
		return -1, err
	}
	if count == 0 {
		count, err = db.FollowerCount(s.ctx, req.UserId)
	}
	if err != nil {
		return -1, nil
	}
	return count, nil

}
