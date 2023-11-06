package service

import (
	"errors"
	"github.com/benxinm/tiktok/cmd/follow/dal/cache"
	"github.com/benxinm/tiktok/cmd/follow/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/follow"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

func (s *FollowService) IsFollow(req *follow.IsFollowRequest) (bool, error) {
	isFollow, err := cache.IsFollow(s.ctx, req.UserId, req.ToUserId)
	if err != nil {
		return false, err
	}
	if isFollow {
		return true, nil
	}
	isFollowDb, err := db.IsFollow(s.ctx, req.UserId, req.ToUserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		klog.Errorf("db is follow error: %v\n", err)
		return false, err
	}
	return isFollowDb, nil
}
