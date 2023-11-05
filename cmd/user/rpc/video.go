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

func GetWorkCount(ctx context.Context, req *video.GetWorkCountRequest) (int64, error) {
	resp, err := videoClient.GetWorkCount(ctx, req)
	if err != nil {
		return -1, err
	}
	if resp.Base.Code != myerrors.SuccessCode {
		return -1, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}
	return resp.WorkCount, nil
}
