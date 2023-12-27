package main

import (
	"context"
	"github.com/benxinm/tiktok/cmd/chat/pack"
	"github.com/benxinm/tiktok/cmd/chat/service"
	chat "github.com/benxinm/tiktok/kitex_gen/chat"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"github.com/benxinm/tiktok/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessagePost implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessagePost(ctx context.Context, req *chat.MessagePostRequest) (resp *chat.MessagePostReponse, err error) {
	resp = new(chat.MessagePostReponse)
	claim, err := utils.VerifyToken(req.Token)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, err
	}

	err = service.NewChatService(ctx).SendMessage(req, claim.UserId, time.Now().Format(time.RFC3339))
	if err != nil {
		klog.Error(err)
		resp.Base = pack.MakeBaseResp(err)
		return resp, err
	}
	resp.Base = pack.MakeBaseResp(nil)
	return
}

// MessageList implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageList(ctx context.Context, req *chat.MessageListRequest) (resp *chat.MessageListResponse, err error) {
	resp = new(chat.MessageListResponse)
	claim, err := utils.VerifyToken(req.Token)
	if err != nil || claim == nil {
		klog.Error(err)
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, err
	}
	// 获取消息列表
	// redis中存在则返回，不存在查询mysql,
	messageList, err := service.NewChatService(ctx).GetMessages(req, claim.UserId)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.MakeBaseResp(err)
		resp.MessageList = pack.MakeMessages(nil)
		resp.Total = 0
		return resp, err
	}
	resp.Base = pack.MakeBaseResp(nil)
	resp.MessageList = pack.MakeMessages(messageList)
	resp.Total = int64(len(messageList))
	return
}
