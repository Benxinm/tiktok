package service

import (
	"errors"
	"github.com/benxinm/tiktok/cmd/chat/dal/db"
	"github.com/benxinm/tiktok/cmd/chat/dal/mq"
	"github.com/benxinm/tiktok/kitex_gen/chat"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
)

func (c *ChatService) SendMessage(req *chat.MessagePostRequest, user_id int64, create_at string) error {
	if len(req.Content) == 0 || len(req.Content) > 1000 {
		klog.Error("character limit error")
		return errors.New("")
	}
	message := &mq.MiddleMessage{
		Id:         db.SF.NextVal(),
		ToUserId:   req.ToUserId,
		FromUserId: user_id,
		IsReadNum:  make([]int64, 0),
		Content:    req.Content,
		CreatedAt:  create_at,
	}
	mid, err := sonic.Marshal(message)
	if err != nil {
		klog.Error(err)
		return err
	}
	err = mq.ChatMQCli.Publish(c.ctx, string(mid))
	if err != nil {
		klog.Error(err)
		return err
	}
	return nil
}
