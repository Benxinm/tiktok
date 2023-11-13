package service

import (
	"github.com/benxinm/tiktok/cmd/interaction/dal/cache"
	"github.com/benxinm/tiktok/cmd/interaction/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) GetCommentCount(req *interaction.CommentCountRequest) (count int64, err error) {
	commentCount, err := cache.GetCommentCount(s.ctx, req.VideoId)
	if err != nil {
		return 0, err
	}
	if commentCount == -1 {
		commentCount, err = db.CountCommentsByVideoId(s.ctx, req.VideoId)
		if err != nil {
			return 0, err
		}
		cache.SetCommentCount(s.ctx, req.VideoId, commentCount)
	}
	return commentCount, nil
}
