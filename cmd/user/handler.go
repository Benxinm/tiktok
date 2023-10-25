package main

import (
	"context"
	"github.com/benxinm/tiktok/cmd/user/pack"
	"github.com/benxinm/tiktok/cmd/user/service"
	user "github.com/benxinm/tiktok/kitex_gen/user"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"github.com/benxinm/tiktok/pkg/utils"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp = new(user.RegisterResponse)

	if len(req.Username) == 0 || len(req.Username) > 255 || len(req.Password) == 0 || len(req.Password) > 255 {
		resp.Base = pack.MakeBaseResp(myerrors.ParamError)
		return resp, nil
	}
	userService := service.NewUserService(ctx)
	userResp, err := userService.CreateUser(req)

	if err != nil {
		resp.Base = pack.MakeBaseResp(err)
		return resp, err
	}
	token, err := utils.GenToken(userResp.Id)
	if err != nil {
		resp.Base = pack.MakeBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.MakeBaseResp(nil)
	resp.UserId = userResp.Id
	resp.Token = token
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)
	if len(req.Username) == 0 || len(req.Username) > 255 || len(req.Password) == 0 || len(req.Password) > 255 {
		resp.Base = pack.MakeBaseResp(myerrors.ParamError)
		return resp, nil
	}

	userService := service.NewUserService(ctx)
	userResp, err := userService.UserLogin(req)

	if err != nil {
		resp.Base = pack.MakeBaseResp(err)
		return resp, nil
	}
	token, err := utils.GenToken(userResp.Id)
	if err != nil {
		resp.Base = pack.MakeBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.MakeBaseResp(nil)
	resp.User = pack.User(userResp)
	resp.Token = token
	return
}

// Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) Info(ctx context.Context, req *user.InfoRequest) (resp *user.InfoResponse, err error) {
	resp = new(user.InfoResponse)

	if req.UserId < 10000 {
		resp.Base = pack.MakeBaseResp(myerrors.ParamError)
		return resp, nil
	}
	if _, err := utils.VerifyToken(req.Token); err != nil {
		resp.Base = pack.MakeBaseResp(myerrors.AuthFailedError)
		return resp, nil
	}
	userService := service.NewUserService(ctx)
	userResp, err := userService.GetUser(req)

	if err != nil {
		resp.Base = pack.MakeBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.MakeBaseResp(nil)
	resp.User = userResp
	return
}
