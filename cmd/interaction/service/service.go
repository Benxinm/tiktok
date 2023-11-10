package service

import "context"

type InteractionService struct {
	ctx context.Context
}

var interactionService *InteractionService

func NewUserService(ctx context.Context) *InteractionService {
	if interactionService == nil {
		interactionService = &InteractionService{
			ctx: ctx,
		}
		return interactionService
	} else {
		interactionService.ctx = ctx
		return interactionService
	}
}
