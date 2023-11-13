package service

import (
	"github.com/benxinm/tiktok/cmd/interaction/dal/cache"
	"github.com/benxinm/tiktok/cmd/interaction/dal/db"
	"github.com/benxinm/tiktok/cmd/interaction/pack"
	"github.com/benxinm/tiktok/cmd/interaction/rpc"
	"github.com/benxinm/tiktok/kitex_gen/interaction"
	"github.com/benxinm/tiktok/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/sync/errgroup"
)

func (s *InteractionService) CreateComment(req *interaction.CommentActionRequest, uid int64) (*interaction.Comment, error) {
	var eg errgroup.Group
	commentModel := &db.Comment{
		VideoId: req.VideoId,
		UserId:  uid,
		Content: req.CommentText,
	}
	comment := new(db.Comment)

	eg.Go(func() error {
		defer func() {
			if e := recover(); e != nil {
				klog.Error(e)
			}
		}()
		var err error
		comment, err = db.CreateComment(s.ctx, commentModel)
		return err
	})
	eg.Go(func() error {
		defer func() {
			if e := recover(); e != nil {
				klog.Error(e)
			}
		}()
		count, err := cache.GetCommentCount(s.ctx, req.VideoId)
		if err != nil {
			return err
		}
		if count >= 0 {
			err = cache.AddCommentCount(s.ctx, req.VideoId)
		}
		return err
	})

	userInfo := new(user.User)

	eg.Go(func() error {
		defer func() {
			if e := recover(); e != nil {
				klog.Error(e)
			}
		}()
		var err error
		userInfo, err = rpc.GetUser(s.ctx, &user.InfoRequest{
			UserId: uid,
			Token:  req.Token,
		})
		return err
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return pack.Comment(comment, userInfo), nil
}
