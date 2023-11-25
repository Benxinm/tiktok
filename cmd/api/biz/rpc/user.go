package rpc

import (
	"context"
	"github.com/benxinm/tiktok/kitex_gen/user"
	"github.com/benxinm/tiktok/kitex_gen/user/userservice"
	"github.com/benxinm/tiktok/pkg/constants"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func InitUserRpc() {
	//TODO 改为config获取etcd地址
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:7890"})
	if err != nil {
		panic(err)
	}
	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectionTimeout),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func UserRegister(ctx context.Context, req *user.RegisterRequest) (int64, string, error) {
	resp, err := userClient.Register(ctx, req)
	if err != nil {
		return -1, "", err
	}
	if resp.Base.Code != myerrors.SuccessCode {
		return -1, "", myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}
	return resp.UserId, resp.Token, nil
}

func UserLogin(ctx context.Context, req *user.LoginRequest) (int64, string, error) {
	resp, err := userClient.Login(ctx, req)

	if err != nil {
		return -1, "", err
	}

	if resp.Base.Code != myerrors.SuccessCode {
		return -1, "", myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}

	return resp.User.Id, resp.Token, nil
}

func UserInfo(ctx context.Context, req *user.InfoRequest) (*user.User, error) {
	resp, err := userClient.Info(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != myerrors.SuccessCode {
		return nil, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}

	return resp.User, nil
}
