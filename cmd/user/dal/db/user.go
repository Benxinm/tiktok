package db

import (
	"context"
	"errors"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id              int64
	Username        string
	Password        string
	Avatar          string `gorm:"default:https://my-tiktok-video.oss-cn-shanghai.aliyuncs.com/default_avatar.jpg?Expires=1697717543&OSSAccessKeyId=TMP.3Kj1NgdKEbvNCXXwAKF6ETfZpGFEqQ27JrqvLRyGkZbnQzatdSGUkaHh4DaGDLXUPKW8X2YhEMfHCwwKzaZSy6uxgTVGoh&Signature=MlkawWxCVp43XleUAzRDRBrItWc%3D"`
	BackgroundImage string `gorm:"default:https://my-tiktok-video.oss-cn-shanghai.aliyuncs.com/default_backimg.jpg?Expires=1697717631&OSSAccessKeyId=TMP.3Kj1NgdKEbvNCXXwAKF6ETfZpGFEqQ27JrqvLRyGkZbnQzatdSGUkaHh4DaGDLXUPKW8X2YhEMfHCwwKzaZSy6uxgTVGoh&Signature=H%2Bgw3TFVsjzrMhtVOFiHJfcSerw%3D"`
	Signature       string `gorm:"default:快留下你的个签吧~"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func CreateUser(ctx context.Context, user *User) (*User, error) {
	userResp := new(User)
	err := DB.WithContext(ctx).Where("username = ?", user.Username).First(&userResp).Error

	if err == nil {
		return nil, myerrors.UserExistedError
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err := DB.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByUsername(ctx context.Context, username string) (*User, error) {
	userResp := new(User)
	err := DB.WithContext(ctx).Where("username = ?", username).First(&userResp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerrors.UserNotFoundError
		}
		return nil, err
	}
	return userResp, nil
}

func GetUserByID(ctx context.Context, id int64) (*User, error) {
	userResp := new(User)
	err := DB.WithContext(ctx).Where("id = ?", id).First(&userResp).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerrors.UserNotFoundError
		}
		return nil, err // 加个未知错误
	}
	return userResp, nil
}
