package rpc

import (
	"context"
	"github.com/benxinm/tiktok/kitex_gen/video"
	"github.com/benxinm/tiktok/kitex_gen/video/videoservice"
	"github.com/benxinm/tiktok/pkg/constants"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func InitVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1"}) //TODO config file
	if err != nil {
		panic(err)
	}
	c, err := videoservice.NewClient(
		constants.VideoServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectionTimeout),
		client.WithResolver(r),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		//TODO MORE CONFIG?
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func VideoFeed(ctx context.Context, req *video.FeedRequest) ([]*video.Video, int64, error) {
	resp, err := videoClient.Feed(ctx, req)

	if err != nil {
		return nil, -1, err
	}

	if resp.Base.Code != myerrors.SuccessCode {
		return nil, -1, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}

	return resp.VideoList, resp.NextTime, nil
}

func PublishList(ctx context.Context, req *video.GetPublishListRequest) ([]*video.Video, error) {
	resp, err := videoClient.GetPublishList(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != myerrors.SuccessCode {
		return nil, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}

	return resp.ListVideo, nil
}

func VideoPublish(ctx context.Context, req *video.PutVideoRequest) error {
	resp, err := videoClient.PutVideo(ctx, &video.PutVideoRequest{
		VideoFile: req.VideoFile,
		Title:     req.Title,
		Token:     req.Token,
	})

	if err != nil {
		return err
	}

	if resp.Base.Code != myerrors.SuccessCode {
		return myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}
	return nil
}
