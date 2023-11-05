package service

import (
	"errors"
	"github.com/benxinm/tiktok/cmd/follow/dal/cache"
	"github.com/benxinm/tiktok/cmd/follow/dal/db"
	"github.com/benxinm/tiktok/cmd/follow/pack"
	"github.com/benxinm/tiktok/cmd/follow/rpc"
	"github.com/benxinm/tiktok/kitex_gen/follow"
	"github.com/benxinm/tiktok/kitex_gen/user"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

func (s *FollowService) GetFollowList(req *follow.FollowListRequest) ([]*follow.User, error) {
	userList := make([]*follow.User, 0, 10)
	list, err := cache.FollowList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	} else if len(list) == 0 {
		list, err = db.FollowList(s.ctx, req.UserId)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userList, err
		} else if err != nil {
			return nil, err
		}
		//TODO db => redis
	}
	var eg errgroup.Group
	for index, id := range list {
		eg.Go(func() error {
			user, err := rpc.GetUser(s.ctx, &user.InfoRequest{
				UserId: id,
				Token:  req.Token,
			})
			if err != nil {
				return err
			}
			follow := pack.User(user)
			//线程安全
			userList[index] = follow
			return nil
		})
	}
	if err = eg.Wait(); err != nil {
		return nil, err
	}
	return userList, nil
}
