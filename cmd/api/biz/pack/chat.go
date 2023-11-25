package pack

import (
	"github.com/benxinm/tiktok/cmd/api/biz/model/api"
	"github.com/benxinm/tiktok/kitex_gen/chat"
	"github.com/cloudwego/kitex/pkg/klog"
)

func MessageList(list []*chat.Message) []*api.Message {
	resp := make([]*api.Message, 0)

	var createtime string

	for _, data := range list {
		if data.CreateTime == "" {
			createtime = "0"
		} else {
			createtime = data.CreateTime
		}

		klog.Infof("createtime: %v\n", createtime)
		resp = append(resp, &api.Message{
			ID:         data.Id,
			ToUserID:   data.ToUserId,
			FromUserID: data.FromUserId,
			Content:    data.Content,
			CreateTime: data.CreateTime,
		})
	}

	return resp
}
