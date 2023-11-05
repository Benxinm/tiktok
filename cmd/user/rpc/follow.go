package rpc

import (
	"context"
	"github.com/benxinm/tiktok/kitex_gen/follow"
	"github.com/benxinm/tiktok/kitex_gen/follow/followservice"
	"github.com/benxinm/tiktok/pkg/constants"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func InitFollowRpc() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1"}) //TODO config file
	if err != nil {
		panic(err)
	}
	c, err := followservice.NewClient(
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
	followClient = c
}

func GetFollowCount(ctx context.Context, req *follow.FollowCountRequest) (int64, error) {
	resp, err := followClient.FollowCount(ctx, req)
	if err != nil {
		return -1, err
	}
	if resp.Base.Code != myerrors.SuccessCode {
		return -1, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}
	return resp.FollowCount, nil
}

func GetFollowerCount(ctx context.Context, req *follow.FollowerCountRequest) (int64, error) {
	resp, err := followClient.FollowerCount(ctx, req)
	if err != nil {
		return -1, err
	}
	if resp.Base.Code != myerrors.SuccessCode {
		return -1, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}
	return resp.FollowerCount, nil
}

func IsFollow(ctx context.Context, req *follow.IsFollowRequest) (bool, error) {
	resp, err := followClient.IsFollow(ctx, req)
	if err != nil {
		return false, err
	}
	if resp.Base.Code != myerrors.SuccessCode {
		return false, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}
	return resp.IsFollow, nil
}
