package main

import (
	"context"
	"fmt"
	"github.com/benxinm/tiktok/cmd/video/pack"
	"github.com/benxinm/tiktok/cmd/video/service"
	video "github.com/benxinm/tiktok/kitex_gen/video"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"github.com/benxinm/tiktok/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/sync/errgroup"
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
	resp = new(video.PutVideoResponse)
	claim, err := utils.VerifyToken(req.Token)
	if err != nil {
		return nil, myerrors.AuthFailedError
	}
	videoName := pack.GenVideoName(claim.UserId)
	coverName := pack.GenCoverName(claim.UserId)

	var eg errgroup.Group

	eg.Go(func() error {
		err = service.NewVideoService(ctx).UploadVideo(req, videoName)
		if err != nil {
			klog.Error(err)
			return err
		}
		return nil
	})
	eg.Go(func() error {
		err = service.NewVideoService(ctx).UploadCover(req, coverName)
		if err != nil {
			klog.Error(err)
			return err
		}
		return nil
	})

	eg.Go(func() error { //TODO endpoint should get from config file
		playUrl := fmt.Sprintf("https://%s/%s/%s", "endpoint", "direction", videoName)
		coverUrl := fmt.Sprintf("https://%s/%s/%s", "endpoint", "direction", coverName)
		_, err = service.NewVideoService(ctx).Create(req, claim.UserId, playUrl, coverUrl)
		if err != nil {
			klog.Error(err)
			return err
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		resp.Base = pack.MakeBaseResp(err)
		return resp, err
	}
	resp.Base = pack.MakeBaseResp(nil)
	return resp, nil
}

// GetFavoriteVideoInfo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFavoriteVideoInfo(ctx context.Context, req *video.GetFavoriteVideoInfoRequest) (resp *video.GetFavoriteVideoInfoResponse, err error) {
	resp = new(video.GetFavoriteVideoInfoResponse)
	if len(req.VideoId) == 0 {
		resp.Base = pack.MakeBaseResp(myerrors.ParamError)
		return resp, nil
	}
	if _, err := utils.VerifyToken(req.Token); err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}

	//videoList,userList/
	return
}

// GetPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishList(ctx context.Context, req *video.GetPublishListRequest) (resp *video.GetPublishListResponse, err error) {
	resp = new(video.GetPublishListResponse)
	if _, err := utils.VerifyToken(req.Token); err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}
	if req.UserId < 10000 {
		resp.Base = pack.MakeBaseResp(myerrors.ParamError)
		return resp, nil
	}
	videoList, err := service.NewVideoService(ctx).GetPublishVideoInfo(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.MakeBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.MakeBaseResp(nil)
	resp.ListVideo = videoList
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
