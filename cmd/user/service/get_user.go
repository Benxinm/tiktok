package service

import (
	"github.com/benxinm/tiktok/cmd/user/dal/db"
	"github.com/benxinm/tiktok/cmd/user/pack"
	"github.com/benxinm/tiktok/kitex_gen/user"
)

func (service *UserService) GetUser(req *user.InfoRequest) (*user.User, error) {
	userResp := new(user.User)
	userModel, err := db.GetUserByID(service.ctx, req.UserId)
	userResp = pack.User(userModel)
	if err != nil {
		return nil, err
	}
	//rpc needed
	return userResp, nil
}
