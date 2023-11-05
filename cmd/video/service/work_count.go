package service

import (
	"github.com/benxinm/tiktok/cmd/video/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/video"
)

func (s *VideoService) WorkCount(req *video.GetWorkCountRequest) (int64, error) {
	return db.GetWorkCountByUid(s.ctx, req.UserId)
}
