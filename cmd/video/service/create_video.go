package service

import (
	"github.com/benxinm/tiktok/cmd/video/dal/db"
	"github.com/benxinm/tiktok/kitex_gen/video"
)

func (s *VideoService) Create(req *video.PutVideoRequest, userId int64, playUrl string, coverUrl string) (*db.Video, error) {
	videoModel := &db.Video{
		UserID:   userId,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    req.Title,
	}
	return db.CreateVideo(s.ctx, videoModel)
}
