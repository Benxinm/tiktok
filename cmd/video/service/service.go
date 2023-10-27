package service

import (
	"context"
)

type VideoService struct {
	ctx context.Context
}

var videoService *VideoService

func NewVideoService(ctx context.Context) *VideoService {
	if videoService == nil {
		videoService = &VideoService{
			ctx: ctx,
		}
		return videoService
	} else {
		videoService.ctx = ctx
		return videoService
	}
	return
}