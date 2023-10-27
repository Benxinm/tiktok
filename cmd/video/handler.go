package main

import (
	"context"
	"github.com/benxinm/tiktok/cmd/video/pack"
	video "github.com/benxinm/tiktok/kitex_gen/video"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"github.com/benxinm/tiktok/pkg/utils"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	resp = new(video.FeedResponse)
	if _, err := utils.VerifyToken(req.Token); err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, err
	}

	return
}

// PutVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PutVideo(ctx context.Context, req *video.PutVideoRequest) (resp *video.PutVideoResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFavoriteVideoInfo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFavoriteVideoInfo(ctx context.Context, req *video.GetFavoriteVideoInfoRequest) (resp *video.GetFavoriteVideoInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// GetPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishList(ctx context.Context, req *video.GetPublishListRequest) (resp *video.GetPublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetWorkCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetWorkCount(ctx context.Context, req *video.GetWorkCountRequest) (resp *video.GetWorkCountResponse, err error) {
	// TODO: Your code here...
	return
}

// GetVideoIDByUid implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoIDByUid(ctx context.Context, req *video.GetVideoIDByUidRequest) (resp *video.GetVideoIDByUidResponse, err error) {
	// TODO: Your code here...
	return
}
