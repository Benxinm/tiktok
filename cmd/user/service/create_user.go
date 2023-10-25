package service

import (
	"github.com/benxinm/tiktok/cmd/user/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/user"
)

func (service *UserService) CreateUser(req *user.RegisterRequest) (*db.User, error) {
	user := &db.User{
		Username: req.Username,
		Password: req.Password,
	}
	return db.CreateUser(service.ctx, user)
}
