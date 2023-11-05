package service

import (
	"github.com/benxinm/tiktok/cmd/follow/dal/cache"
	"github.com/benxinm/tiktok/cmd/follow/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/follow"
)

func (s *FollowService) GetFollowCount(req *follow.FollowCountRequest) (int64, error) {
	followCount, err := cache.FollowCount(s.ctx, req.UserId)
	if err != nil {
		return -1, err
	}
	if followCount == 0 {
		followCount, err = db.FollowCount(s.ctx, req.UserId)
	}
	if err != nil {
		return -1, err
	}
	return followCount, nil
}
