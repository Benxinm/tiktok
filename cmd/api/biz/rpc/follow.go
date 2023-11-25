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

func RelationAction(ctx context.Context, req *follow.ActionRequest) error {
	resp, err := followClient.Action(ctx, req)

	if err != nil {
		return err
	}

	if resp.Base.Code != myerrors.SuccessCode {
		return myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}

func FollowList(ctx context.Context, req *follow.FollowListRequest) ([]*follow.User, error) {
	resp, err := followClient.FollowList(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != myerrors.SuccessCode {
		return nil, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}

	return resp.UserList, nil
}

func FollowerList(ctx context.Context, req *follow.FollowerListRequest) ([]*follow.User, error) {
	resp, err := followClient.FollowerList(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != myerrors.SuccessCode {
		return nil, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}

	return resp.UserList, nil
}

func FriendList(ctx context.Context, req *follow.FriendListRequest) ([]*follow.FriendUser, error) {
	resp, err := followClient.FriendList(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != myerrors.SuccessCode {
		return nil, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}

	return resp.UserList, nil
}
