package service

import "context"

type ChatService struct {
	ctx context.Context
}

var chatService *ChatService

func NewChatService(ctx context.Context) *ChatService {
	if chatService == nil {
		chatService = &ChatService{
			ctx: ctx,
		}
		return chatService
	} else {
		chatService.ctx = ctx
		return chatService
	}
}
