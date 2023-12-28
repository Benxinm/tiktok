package db

import (
	"context"
	"errors"
	"github.com/benxinm/tiktok/pkg/constants"
	"gorm.io/gorm"
	"time"
)

type Favorite struct {
	ID        int64 `json:"id"`
	UserID    int64 `json:"user_id"`
	VideoID   int64 `json:"video_id"`
	Status    int64 `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type VideoFavourite struct {
	ID            int64 `json:"id"`
	VideoID       int64 `json:"video_id"`
	FavoriteCount int64 `json:"favorite_count"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func IsFavorited(ctx context.Context, uid int64, vid int64, status int64) error {
	var fav Favorite
	return DB.Table(constants.FavoriteTableName).WithContext(ctx).
		Where("user_id = ? AND video_id = ? AND status = ?", uid, vid, status).First(&fav).Error
}

func IsFavoriteExist(ctx context.Context, uid int64, vid int64) (bool, error) {
	var resp Favorite
	err := DB.Table(constants.FavoriteTableName).WithContext(ctx).Where("user_id = ? and video_id = ? and status = 1", uid, vid).First(&resp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func CreateFavorite(ctx context.Context, favorite *Favorite) error {
	favorite.ID = SF.NextVal()
	return DB.Table(constants.FavoriteTableName).WithContext(ctx).Create(favorite).Error
}
func UpdateFavoriteStatus(ctx context.Context, uid int64, vid int64, status int64) error {
	return DB.Table(constants.FavoriteTableName).WithContext(ctx).
		Where("user_id = ? AND video_id = ?", uid, vid).Update("status", status).Error
}

func GetFavouriteVideosByUid(ctx context.Context, uid int64) ([]int64, error) {
	var vids []int64
	err := DB.Table(constants.FavoriteTableName).WithContext(ctx).Select("video_id").Where("user_id = ? and status = 1", uid).Find(&vids).Error
	if err != nil {
		return nil, err
	}
	return vids, nil
}

func GetVideoFavouriteCount(ctx context.Context, vid int64) (int64, error) {
	var vf VideoFavourite
	if err := DB.Table(constants.VideoFavouriteTableName).WithContext(ctx).Where("video_id = ?", vid).First(vf).Error; err != nil {
		return 0, err
	}
	return vf.FavoriteCount, nil
}

func GetUserFavouriteCount(ctx context.Context, uid int64) (int64, error) {
	var count int64
	if err := DB.Table(constants.FavoriteTableName).WithContext(ctx).
		Where("user_id = ? AND status = 1", uid).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func UpdateVideoFavouriteCount(ctx context.Context, vid int64, count int64) error {
	return DB.Table(constants.VideoFavouriteTableName).WithContext(ctx).Where("video_id = ?", vid).Update("favourite_count", count).Error
}
