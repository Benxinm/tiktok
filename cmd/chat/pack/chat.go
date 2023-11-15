package pack

import (
	"fmt"
	"github.com/benxinm/tiktok/cmd/chat/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/chat"
)

func MakeMessages(data []*db.Message) []*chat.Message {
	if data == nil {
		return nil
	}
	messages := make([]*chat.Message, 0, len(data))
	for _, message := range data {
		msg := &chat.Message{
			Id:         message.Id,
			ToUserId:   message.ToUserId,
			FromUserId: message.FromUserId,
			Content:    message.Content,
			CreateTime: fmt.Sprintf("%v", message.CreatedAt.UnixMilli()),
		}
		messages = append(messages, msg)
	}
	return messages
}
