package db

import (
	"context"
	"errors"
	"github.com/benxinm/tiktok/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	Id         int64  `json:"id"`
	ToUserId   int64  `json:"to_user_id"`
	FromUserId int64  `json:"from_user_id"`
	Content    string `json:"content"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type MiddleMessage struct {
	Id         int64
	ToUserId   int64
	FromUserId int64
	Content    string
	IsReadNum  []int64
	CreatedAt  string
}

type MessageArray []*Message

func GetMessageList(ctx context.Context, tuid int64, fuid int64) ([]*Message, error) {
	var messages []*Message
	err := DB.WithContext(ctx).
		Select("id", "to_user_id", "from_user_id", "content", "created_at").
		Where("(to_user_id=? AND from_user_id =?) OR (to_user_id=? AND from_user_id =?) ", tuid, fuid, fuid, tuid).
		Order("created_at desc").
		Find(&messages).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		klog.Errorf("messages error: %v\n", err)
		return nil, err
	}
	return messages, nil
}

func CreateMessage(message *Message) error {
	return DB.Table(constants.MessageTableName).Create(message).Error
}

func Convert(message *Message, tempMessage *MiddleMessage) (err error) {
	message.Id = tempMessage.Id
	message.ToUserId = tempMessage.ToUserId
	message.FromUserId = tempMessage.FromUserId
	message.Content = tempMessage.Content
	message.CreatedAt, err = time.Parse(time.RFC3339, tempMessage.CreatedAt)
	if err != nil {
		return err
	}
	return
}

//  implement the interface

func (array MessageArray) Len() int {
	return len(array)
}

func (array MessageArray) Less(i, j int) bool {
	return array[i].CreatedAt.Unix() > array[j].CreatedAt.Unix()
}

func (array MessageArray) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]
}
