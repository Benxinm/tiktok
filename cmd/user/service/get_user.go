package service

import (
	"github.com/benxinm/tiktok/cmd/user/dal/db"
	"github.com/benxinm/tiktok/cmd/user/pack"
	"github.com/benxinm/tiktok/cmd/user/rpc"
	"github.com/benxinm/tiktok/kitex_gen/follow"
	"github.com/benxinm/tiktok/kitex_gen/user"
	"github.com/benxinm/tiktok/kitex_gen/video"
	"github.com/benxinm/tiktok/pkg/utils"
)

func (s *UserService) GetUser(req *user.InfoRequest) (*user.User, error) {
	userResp := new(user.User)
	userModel, err := db.GetUserByID(s.ctx, req.UserId)
	userResp = pack.User(userModel)
	if err != nil {
		return nil, err
	}
	claim, err := utils.VerifyToken(req.Token)
	if err != nil {
		return nil, err
	}
	//rpc needed
	count, err := rpc.GetFollowCount(s.ctx, &follow.FollowCountRequest{
		UserId: userModel.Id, Token: req.Token,
	})
	if err != nil {
		return nil, err
	}
	userResp.FollowCount = count
	followerCount, err := rpc.GetFollowerCount(s.ctx, &follow.FollowerCountRequest{
		UserId: userModel.Id, Token: req.Token,
	})
	if err != nil {
		return nil, err
	}
	userResp.FollowerCount = followerCount
	isFollow, err := rpc.IsFollow(s.ctx, &follow.IsFollowRequest{UserId: claim.UserId, ToUserId: req.UserId, Token: req.Token})
	if err != nil {
		return nil, err
	}
	userResp.IsFollow = isFollow
	workCount, err := rpc.GetWorkCount(s.ctx, &video.GetWorkCountRequest{UserId: userModel.Id, Token: req.Token})
	if err != nil {
		return nil, err
	}
	userResp.WorkCount = workCount
	//TODO Favourite video list needed
	return userResp, nil
}
