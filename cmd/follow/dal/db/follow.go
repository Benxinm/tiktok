package db

import (
	"context"
	"errors"
	"github.com/benxinm/tiktok/pkg/constants"
	"github.com/benxinm/tiktok/pkg/myerrors"
	"gorm.io/gorm"
	"time"
)

type Follow struct {
	Id        int64
	UserId    int64
	ToUserId  int64
	Status    int64 //保证数据不进行删除 0-关注 1-取关
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func IsFollow(ctx context.Context, uid int64, tuid int64) (bool, error) {
	var model *Follow
	err := DB.WithContext(ctx).Model(&Follow{}).Where("user_id= ? AND to_user_id = ?", uid, tuid).First(model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // db中也查不到
			return false, gorm.ErrRecordNotFound
		}
		return false, err
	}
	return model.Status == 1, nil
}

func FollowAction(ctx context.Context, follow *Follow) error {
	resp := new(Follow)
	err := DB.WithContext(ctx).Model(&Follow{}).Where("user_id= ? AND to_user_id = ?", follow.UserId, follow.ToUserId).First(&resp).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		follow.Id = SF.NextVal()
		follow.Status = 0
		return DB.WithContext(ctx).Create(follow).Error
	} else if err != nil {
		return err
	}
	err = DB.WithContext(ctx).Model(&Follow{}).Where("user_id= ? AND to_user_id = ?", follow.UserId, follow.ToUserId).Update("status", 0).Error
	if err != nil {
		return err
	}
	return nil
}

func UnFollowAction(ctx context.Context, follow *Follow) error {
	_, err := IsFollow(ctx, follow.UserId, follow.ToUserId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return myerrors.FollowNotFoundError
	} else if err != nil {
		return err
	}
	err = DB.WithContext(ctx).Model(&Follow{}).Where("user_id= ? AND to_user_id = ?", follow.UserId, follow.ToUserId).Update("status", 1).Error
	if err != nil {
		return err
	}
	return nil
}

func FollowList(ctx context.Context, uid int64) ([]int64, error) {
	var followList []int64

	err := DB.WithContext(ctx).Table(constants.FollowTableName).Select("to_user_id").Where("user_id = ? AND status = ?", uid, 0).Find(&followList).Error
	if err != nil {
		return nil, err
	}

	if len(followList) == 0 { // db中也查不到
		return nil, gorm.ErrRecordNotFound
	}

	return followList, nil
}

func FollowerList(ctx context.Context, uid int64) (*[]int64, error) {
	var followerList []int64

	err := DB.WithContext(ctx).Table(constants.FollowTableName).Select("user_id").Where("to_user_id = ? AND status = ?", uid, 0).Find(&followerList).Error
	if err != nil {
		return nil, err
	}

	if len(followerList) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &followerList, nil
}

//TODO FriendList idea

func FollowCount(ctx context.Context, uid int64) (int64, error) {
	var count int64
	err := DB.WithContext(ctx).Model(&Follow{}).Where("user_id = ? AND status = ?", uid, 0).Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}

func FollowerCount(ctx context.Context, uid int64) (int64, error) {
	var count int64
	err := DB.WithContext(ctx).Model(&Follow{}).Where("to_user_id = ? AND status = ?", uid, 0).Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}
