package db

import (
	"context"
	"github.com/benxinm/tiktok/pkg/constants"
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	VideoId   int64  `json:"video_id"`
	Content   string `json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" msg:"-"`
}

func CreateComment(ctx context.Context, comment *Comment) (*Comment, error) {
	if err := DB.Table(constants.CommentTableName).WithContext(ctx).Create(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func DeleteComment(ctx context.Context, comment *Comment) (*Comment, error) {
	if err := DB.Table(constants.CommentTableName).WithContext(ctx).Delete(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func GetCommentById(ctx context.Context, commentId int64) (*Comment, error) {
	resp := new(Comment)

	err := DB.Table(constants.CommentTableName).WithContext(ctx).Where("id = ?", commentId).First(resp).Error
	if err != nil {
		return nil, err
	}
	return resp, err
}

func GetCommentsByVideoId(ctx context.Context, videoId int64) ([]Comment, error) {
	var resp []Comment
	err := DB.Table(constants.CommentTableName).WithContext(ctx).Where("video_id = ? and deleted_at IS NULL", videoId).Order("created_at desc").Find(&resp).Error
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func CountCommentsByVideoId(ctx context.Context, videoId int64) (int64, error) {
	var count int64
	err := DB.Table(constants.CommentTableName).WithContext(ctx).Where("video_id = ? and deleted_at IS NULL").Count(&count).Error
	if err != nil {
		return -1, err
	}
	return count, nil
}


