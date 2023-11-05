package service

import "context"

type FollowService struct {
	ctx context.Context
}

var followService *FollowService

func NewFollowService(ctx context.Context) *FollowService {
	if followService == nil {
		followService = &FollowService{
			ctx: ctx,
		}
		return followService
	} else {
		followService.ctx = ctx
		return followService
	}
}
