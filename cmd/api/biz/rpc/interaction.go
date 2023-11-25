package rpc

import (
	"context"
	"github.com/benxinm/tiktok/kitex_gen/interaction"
	"github.com/benxinm/tiktok/kitex_gen/interaction/interactionservice"
	"github.com/benxinm/tiktok/pkg/constants"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func InitInteractionRpc() {
	//TODO 改为config获取etcd地址
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:7890"})
	if err != nil {
		panic(err)
	}
	c, err := interactionservice.NewClient(
		constants.UserServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectionTimeout),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	interactionClient = c
}

func FavoriteAction(ctx context.Context, req *interaction.FavoriteActionRequest) error {
	resp, err := interactionClient.FavoriteAction(ctx, req)

	if err != nil {
		return err
	}

	if resp.Base.Code != myerrors.SuccessCode {
		return myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}

func FavoriteList(ctx context.Context, req *interaction.FavoriteListRequest) ([]*interaction.Video, error) {
	resp, err := interactionClient.FavoriteList(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != myerrors.SuccessCode {
		return nil, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}

	return resp.VideoList, nil
}

func CommentAction(ctx context.Context, req *interaction.CommentActionRequest) (*interaction.Comment, error) {
	resp, err := interactionClient.CommentAction(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != myerrors.SuccessCode {
		return nil, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}

	return resp.Comment, nil
}

func CommentList(ctx context.Context, req *interaction.CommentListRequest) ([]*interaction.Comment, error) {
	resp, err := interactionClient.CommentList(ctx, req)

	if err != nil {
		return nil, err
	}

	if resp.Base.Code != myerrors.SuccessCode {
		return nil, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}

	return resp.CommentList, nil
}
