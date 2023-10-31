package service

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type VideoService struct {
	ctx    context.Context
	bucket *oss.Bucket
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
}
