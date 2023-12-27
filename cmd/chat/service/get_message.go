package service

import (
	"context"
	"github.com/benxinm/tiktok/cmd/chat/dal/cache"
	"github.com/benxinm/tiktok/cmd/chat/dal/db"
	"github.com/benxinm/tiktok/cmd/chat/dal/mq"
	"github.com/benxinm/tiktok/kitex_gen/chat"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/sync/errgroup"
	"sort"
	"strconv"
	"time"
)

func (c *ChatService) GetMessages(req *chat.MessageListRequest, user_id int64) ([]*db.Message, error) {
	mq.Mu.Lock()
	defer mq.Mu.Unlock()
	messageList := make(db.MessageArray, 0)
	key := strconv.FormatInt(req.ToUserId, 10) + "-" + strconv.FormatInt(user_id, 10)
	revkey := strconv.FormatInt(user_id, 10) + "-" + strconv.FormatInt(req.ToUserId, 10)
	no_empty := 0
	msg_array, err := cacheMessageDeal(c.ctx, key, revkey, &no_empty, user_id)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	messageList = append(messageList, msg_array...)
	rev_msg_array, err := cacheMessageDeal(c.ctx, revkey, key, &no_empty, user_id)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	messageList = append(messageList, rev_msg_array...)
	if len(messageList) > 0 {
		sort.Sort(messageList)
		return messageList, nil
	} else if no_empty > 0 {
		return nil, nil
	}
	messages, err := db.GetMessageList(c.ctx, req.ToUserId, user_id)
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	mq_message, err := sonic.Marshal(messages)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	message := make([]*mq.MiddleMessage, 0)
	err = sonic.Unmarshal(mq_message, &message)
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	for _, val := range message {
		val.IsReadNum = append(val.IsReadNum, user_id)
		mes, _ := sonic.Marshal(val)
		key := strconv.FormatInt(val.FromUserId, 10) + "-" + strconv.FormatInt(val.ToUserId, 10)
		cre_time, _ := time.ParseInLocation(time.RFC3339, val.CreatedAt, time.Local)

		err := cache.CreateMessage(c.ctx, key, revkey, cre_time.UnixMilli(), string(mes))
		if err != nil {
			klog.Info(err)
			continue
		}
	}
	return messages, nil
}

func cacheMessageDeal(ctx context.Context, key string, revkey string, isempty *int, user_id int64) (db.MessageArray, error) {
	msg_array := make(db.MessageArray, 0)
	if ok := cache.MessageExist(ctx, key); ok != 0 {
		// 查询 a->b的消息
		mem, err := cache.GetMessage(ctx, key)
		if err != nil {
			klog.Error(err)
			return nil, err
		}
		if len(mem) != 0 {
			*isempty = 1
		}
		var eg errgroup.Group
		for _, val := range mem {
			temp_val := val
			eg.Go(func() error {
				tempMessage := new(db.MiddleMessage)
				message := new(db.Message)
				err = sonic.Unmarshal([]byte(temp_val), &tempMessage)
				if err != nil {
					klog.Error(err)
					return err
				}
				for _, id := range tempMessage.IsReadNum {
					if id == 0 {
						continue
					} else if id == user_id {
						return nil
					}
				}
				err = db.Convert(message, tempMessage)
				if err != nil {
					klog.Error(err)
					return err
				}
				tempMessage.IsReadNum = append(tempMessage.IsReadNum, user_id)
				redis_msg, err := sonic.Marshal(tempMessage)
				if err != nil {
					klog.Error(err)
					return err
				}
				cTime, err := time.ParseInLocation(time.RFC3339, tempMessage.CreatedAt, time.Local)
				if err != nil {
					klog.Error(err)
					return err
				}
				err = cache.CreateMessage(ctx, key, revkey, cTime.UnixMilli(), string(redis_msg))
				if err != nil {
					klog.Error(err)
					return err
				}
				msg_array = append(msg_array, message)
				return nil
			})
		}
		if err := eg.Wait(); err != nil {
			return nil, err
		}
	}
	return msg_array, nil
}
