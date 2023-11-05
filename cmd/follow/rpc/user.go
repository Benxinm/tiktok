package rpc

import (
	"context"
	"github.com/benxinm/tiktok/kitex_gen/user"
	"github.com/benxinm/tiktok/kitex_gen/user/userservice"
	"github.com/benxinm/tiktok/pkg/constants"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func InitUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1"}) //TODO config file
	if err != nil {
		panic(err)
	}
	c, err := userservice.NewClient(
		constants.UserServiceName,
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
	userClient = c
}

func GetUser(ctx context.Context, req *user.InfoRequest) (*user.User, error) {
	resp, err := userClient.Info(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Base.Code != myerrors.SuccessCode {
		return nil, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}
	return resp.User, nil
}
