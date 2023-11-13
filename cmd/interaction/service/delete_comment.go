package service

import (
	"github.com/benxinm/tiktok/cmd/interaction/dal/db"
	"github.com/benxinm/tiktok/cmd/interaction/pack"
	"github.com/benxinm/tiktok/cmd/video/rpc"
	"github.com/benxinm/tiktok/kitex_gen/interaction"
	"github.com/benxinm/tiktok/kitex_gen/user"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"golang.org/x/sync/errgroup"
)

func (s *InteractionService) DeleteComment(req *interaction.CommentActionRequest, uid int64) (*interaction.Comment, error) {
	var eg errgroup.Group
	comment, err := db.GetCommentById(s.ctx, req.CommentId)
	if err != nil {
		return nil, err
	}
	if comment.UserId != uid {
		return nil, myerrors.ParamError
	}
	eg.Go(func() error {
		var err error
		comment, err = db.DeleteComment(s.ctx, comment)
		return err
	})
	var userInfo *user.User
	eg.Go(func() error {
		var err error
		userInfo, err = rpc.GetUser(s.ctx, &user.InfoRequest{
			UserId: uid,
			Token:  req.Token,
		})
		return err
	})
	if err = eg.Wait(); err != nil {
		return nil, err
	}
	return pack.Comment(comment, userInfo), nil
}
