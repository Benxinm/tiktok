package rpc

import (
	"context"
	"github.com/benxinm/tiktok/kitex_gen/chat"
	"github.com/benxinm/tiktok/kitex_gen/chat/messageservice"
	"github.com/benxinm/tiktok/pkg/constants"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func InitChatRpc() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1"}) //TODO config file
	if err != nil {
		panic(err)
	}
	c, err := messageservice.NewClient(
		constants.ChatServiceName,
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
	chatClient = c
}

func MessageAction(ctx context.Context, req *chat.MessagePostRequest) error {
	resp, err := chatClient.MessagePost(ctx, req)

	if err != nil {
		return err
	}

	if resp.Base.Code != myerrors.SuccessCode {
		return myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}

	return nil
}

func MessageList(ctx context.Context, req *chat.MessageListRequest) ([]*chat.Message, int64, error) {
	resp, err := chatClient.MessageList(ctx, req)

	if err != nil {
		return nil, -1, err
	}

	if resp.Base.Code != myerrors.SuccessCode {
		return nil, -1, myerrors.NewMyError(resp.Base.Code, resp.Base.Msg)
	}

	return resp.MessageList, resp.Total, nil
}
