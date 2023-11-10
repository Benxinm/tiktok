package service

import (
	"github.com/benxinm/tiktok/cmd/interaction/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/interaction"
)

func (s *InteractionService) GetCommentCount(req *interaction.CommentCountRequest, time int) (count int64, err error) {
	return db.CountCommentsByVideoId(s.ctx, req.VideoId)
}
