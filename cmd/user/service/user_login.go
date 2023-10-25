package service

import (
	"github.com/benxinm/tiktok/cmd/user/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/user"
	"github.com/benxinm/tiktok/pkg/myerrors"
)

func (service *UserService) UserLogin(req *user.LoginRequest) (*db.User, error) {
	userModel, err := db.GetUserByUsername(service.ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if req.Password != userModel.Password {
		return nil, myerrors.ParamError
	}
	return userModel, nil
}
