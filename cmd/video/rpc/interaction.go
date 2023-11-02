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

func GetVideoLikeCount(ctx context.Context, req *interaction.VideoFavoritedCountRequest) (likes int64, err error) {
	resp, err := interactionClient.VideoFavoritedCount(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.Base.Code != myerrors.SuccessCode {
		return 0, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}
	return resp.LikeCount, nil
}

func GetCommentCount(ctx context.Context, req *interaction.CommentCountRequest) (commentCount int64, err error) {
	resp, err := interactionClient.CommentCount(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.Base.Code != myerrors.SuccessCode {
		return 0, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}
	return resp.CommentCount, nil
}

func GetVideoIsLiked(ctx context.Context, req *interaction.InteractionServiceIsFavoriteArgs) (isFavourite bool, err error) {
	resp, err := interactionClient.IsFavorite(ctx, req.Req)
	if err != nil {
		return false, err
	}
	if resp.Base.Code != myerrors.SuccessCode {
		return false, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}
	return resp.IsFavorite, nil
}
