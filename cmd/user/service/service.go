package service

import "context"

type UserService struct {
	ctx context.Context
}

var userService *UserService

func NewUserService(ctx context.Context) *UserService {
	if userService == nil {
		userService = &UserService{
			ctx: ctx,
		}
		return userService
	} else {
		userService.ctx = ctx
		return userService
	}
}
