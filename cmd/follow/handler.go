package main

import (
	"context"
	follow "github.com/benxinm/tiktok/kitex_gen/follow"
)

// FollowServiceImpl implements the last service interface defined in the IDL.
type FollowServiceImpl struct{}

// Action implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) Action(ctx context.Context, req *follow.ActionRequest) (resp *follow.ActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FollowList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowList(ctx context.Context, req *follow.FollowListRequest) (resp *follow.FollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// FollowerList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowerList(ctx context.Context, req *follow.FollowerListRequest) (resp *follow.FollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// FriendList implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FriendList(ctx context.Context, req *follow.FriendListRequest) (resp *follow.FriendListResponse, err error) {
	// TODO: Your code here...
	return
}

// FollowCount implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowCount(ctx context.Context, req *follow.FollowCountRequest) (resp *follow.FollowCountResponse, err error) {
	// TODO: Your code here...
	return
}

// FollowerCount implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) FollowerCount(ctx context.Context, req *follow.FollowerCountRequest) (resp *follow.FollowerCountResponse, err error) {
	// TODO: Your code here...
	return
}

// IsFollow implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) IsFollow(ctx context.Context, req *follow.IsFollowRequest) (resp *follow.IsFollowResponse, err error) {
	// TODO: Your code here...
	return
}
