package service

import (
	"errors"
	"github.com/benxinm/tiktok/cmd/interaction/dal/db"
	"github.com/benxinm/tiktok/cmd/interaction/pack"
	"github.com/benxinm/tiktok/cmd/interaction/rpc"
	"github.com/benxinm/tiktok/kitex_gen/interaction"
	"github.com/benxinm/tiktok/kitex_gen/user"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

func (s *InteractionService) GetComments(req *interaction.CommentListRequest) ([]*interaction.Comment, error) {
	comments, err := db.GetCommentsByVideoId(s.ctx, req.VideoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, myerrors.CommentNotFoundError
	}
	if err != nil {
		return nil, err
	}
	var eg errgroup.Group
	commentList := make([]*interaction.Comment, len(comments))
	for i := 0; i < len(comments); i++ {
		comment := comments[i]
		commentIndex := i
		eg.Go(func() error {
			userInfo, err := rpc.GetUser(s.ctx, &user.InfoRequest{
				UserId: comment.UserId,
				Token:  req.Token,
			})
			commentList[commentIndex] = pack.Comment(&comment, userInfo)
			return err
		})
	}
	if err = eg.Wait(); err != nil {
		return nil, err
	}
	return commentList, nil
}
