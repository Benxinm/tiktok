package db

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Video struct {
	Id        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	PlayUrl   string `json:"play_url"`
	CoverUrl  string `json:"cover_url"`
	Title     string `json:"title"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func CreateVideo(ctx context.Context, video *Video) (*Video, error) {
	//video.Id = SF.NextVal()
	if err := DB.WithContext(ctx).Create(video).Error; err != nil {
		return nil, err
	}
	return video, nil
}

func GetVideosByIDs(ctx context.Context, ids []int64) ([]Video, error) {
	var videoResp []Video
	if err := DB.WithContext(ctx).Where("id IN ?", ids).Order("created_at").Limit(30).Find(&videoResp).Error; err != nil {
		return nil, err
	}
	return videoResp, nil
}

func GetVideoByTime(ctx context.Context, lastTime string) ([]Video, error) {
	var videoResp []Video
	if err := DB.WithContext(ctx).Where("created_at<?", lastTime).Order("created_at DESC").Limit(30).Find(&videoResp).Error; err != nil {
		return nil, err
	}
	return videoResp, nil
}

func GetVideoByUid(ctx context.Context, uid int64) ([]Video, error) {
	var videoResp []Video
	if err := DB.WithContext(ctx).Where("user_id = ?", uid).Order("created_at DESC").Find(&videoResp).Error; err != nil {
		return nil, err
	}
	return videoResp, nil
}
