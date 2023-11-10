package rpc

import (
	"context"
	"github.com/benxinm/tiktok/kitex_gen/video"
	"github.com/benxinm/tiktok/kitex_gen/video/videoservice"
	"github.com/benxinm/tiktok/pkg/constants"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func InitVideoRpc() {
	//TODO 改为config获取etcd地址
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:7890"})
	if err != nil {
		panic(err)
	}
	c, err := videoservice.NewClient(
		constants.UserServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectionTimeout),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func GetFavoriteVideoList(ctx context.Context, req *video.GetFavoriteVideoInfoRequest) ([]*video.Video, error) {
	resp, err := videoClient.GetFavoriteVideoInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != myerrors.SuccessCode {
		return nil, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}
	return resp.VideoList, nil
}

func GetUserVideoList(ctx context.Context, req *video.GetVideoIDByUidRequest) ([]int64, error) {
	resp, err := videoClient.GetVideoIDByUid(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != myerrors.SuccessCode {
		return nil, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}
	return resp.VideoId, nil
}
