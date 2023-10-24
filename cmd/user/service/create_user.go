package service

import (
	"github.com/benxinm/tiktok/cmd/user/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/user"
)

func (s *UserService) CreateUser(req *user.RegisterRequest) (*db.User, error) {
	user := &db.User{
		Username: req.Username,
		Password: req.Password,
	}
	return db.CreateUser(s.ctx, user)
}
