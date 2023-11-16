package service

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/benxinm/tiktok/cmd/video/dal"
	"github.com/benxinm/tiktok/config"
)

type VideoService struct {
	ctx    context.Context
	bucket *oss.Bucket
}

var videoService *VideoService

func NewVideoService(ctx context.Context) *VideoService {
	if videoService == nil {
		bucket, err := dal.OssClient.Bucket(config.OSS.BucketName)
		if err != nil {
			panic(err)
		}
		videoService = &VideoService{
			ctx:    ctx,
			bucket: bucket,
		}
		return videoService
	} else {
		videoService.ctx = ctx
		return videoService
	}
}
