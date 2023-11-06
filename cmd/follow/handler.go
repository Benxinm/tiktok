package main

import (
	"context"
	"github.com/benxinm/tiktok/cmd/follow/pack"
	"github.com/benxinm/tiktok/cmd/follow/service"
	follow "github.com/benxinm/tiktok/kitex_gen/follow"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"github.com/benxinm/tiktok/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
)

// FollowServiceImpl implements the last service interface defined in the IDL.
type FollowServiceImpl struct{}

// Action implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) Action(ctx context.Context, req *follow.ActionRequest) (resp *follow.ActionResponse, err error) {
	resp = new(follow.ActionResponse)
	if _, err := utils.VerifyToken(req.Token); err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}
	if err := service.NewFollowService(ctx).Follow(req); err != nil {
		klog.Error(err)
		resp.Base = pack.MakeBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.MakeBaseResp(nil)
	return
}

// FollowList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowList(ctx context.Context, req *follow.FollowListRequest) (resp *follow.FollowListResponse, err error) {
	resp = new(follow.FollowListResponse)
	if _, err := utils.VerifyToken(req.Token); err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}
	followResp, err := service.NewFollowService(ctx).GetFollowList(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.MakeBaseResp(err)
	}
	resp.Base = pack.MakeBaseResp(nil)
	resp.UserList = followResp
	return
}

// FollowerList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowerList(ctx context.Context, req *follow.FollowerListRequest) (resp *follow.FollowerListResponse, err error) {
	resp = new(follow.FollowerListResponse)
	if _, err := utils.VerifyToken(req.Token); err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}
	followerResp, err := service.NewFollowService(ctx).GetFollowerList(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.MakeBaseResp(err)
	}
	resp.Base = pack.MakeBaseResp(nil)
	resp.UserList = followerResp
	return
}

// FriendList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FriendList(ctx context.Context, req *follow.FriendListRequest) (resp *follow.FriendListResponse, err error) {
	resp = new(follow.FriendListResponse)
	if _, err := utils.VerifyToken(req.Token); err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}
	// TODO: Your code here...
	return
}

// FollowCount implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowCount(ctx context.Context, req *follow.FollowCountRequest) (resp *follow.FollowCountResponse, err error) {
	resp = new(follow.FollowCountResponse)
	if _, err := utils.VerifyToken(req.Token); err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}
	followResp, err := service.NewFollowService(ctx).GetFollowCount(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.MakeBaseResp(err)
	}
	resp.Base = pack.MakeBaseResp(nil)
	resp.FollowCount = followResp
	return
}

// FollowerCount implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowerCount(ctx context.Context, req *follow.FollowerCountRequest) (resp *follow.FollowerCountResponse, err error) {
	resp = new(follow.FollowerCountResponse)
	if _, err := utils.VerifyToken(req.Token); err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}
	followerResp, err := service.NewFollowService(ctx).GetFollowerCount(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.MakeBaseResp(err)
	}
	resp.Base = pack.MakeBaseResp(nil)
	resp.FollowerCount = followerResp
	return
}

// IsFollow implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) IsFollow(ctx context.Context, req *follow.IsFollowRequest) (resp *follow.IsFollowResponse, err error) {
	resp = new(follow.IsFollowResponse)
	if _, err := utils.VerifyToken(req.Token); err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}
	followResp, err := service.NewFollowService(ctx).IsFollow(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.MakeBaseResp(err)
	}
	resp.Base = pack.MakeBaseResp(nil)
	resp.IsFollow = followResp
	return
}
