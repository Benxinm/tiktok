package service

import (
	"github.com/benxinm/tiktok/cmd/follow/dal/cache"
	"github.com/benxinm/tiktok/cmd/follow/dal/db"
	"github.com/benxinm/tiktok/cmd/video/rpc"
	"github.com/benxinm/tiktok/kitex_gen/follow"
	"github.com/benxinm/tiktok/kitex_gen/user"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"github.com/benxinm/tiktok/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
)

func (s *FollowService) Follow(req *follow.ActionRequest) error {
	claim, err := utils.VerifyToken(req.Token)
	if err != nil {
		return myerrors.AuthFailedError
	}
	if claim.UserId == req.ToUserId {
		return myerrors.ParamError
	}

	_, err = rpc.GetUser(s.ctx, &user.InfoRequest{
		UserId: req.ToUserId,
		Token:  req.Token,
	})
	if err != nil {
		klog.Info(err)
		return myerrors.UserNotFoundError
	}
	acton := db.Follow{UserId: claim.UserId, ToUserId: req.ToUserId}
	switch req.ActionType {
	case 0:
		if err := cache.Follow(s.ctx, acton.UserId, acton.ToUserId); err != nil {
			return err
		}
		if err = db.FollowAction(s.ctx, &acton); err != nil {
			return err
		}
	case 1:
		if err = db.UnFollowAction(s.ctx, &acton); err != nil {
			return err
		}
		if err = cache.UnFollow(s.ctx, acton.UserId, acton.ToUserId); err != nil {
			return err
		}
	default:
		return myerrors.ParamError
	}
	return err
}
