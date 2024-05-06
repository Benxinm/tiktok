package rpc

import (
	"context"
	"github.com/benxinm/tiktok/kitex_gen/follow"
	"github.com/benxinm/tiktok/kitex_gen/follow/followservice"
	"github.com/benxinm/tiktok/pkg/constants"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func InitFollowRpc() {
	//TODO 改为config获取etcd地址
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:7890"})
	if err != nil {
		panic(err)
	}
	c, err := followservice.NewClient(
		constants.UserServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectionTimeout),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	followClient = c
}

func GetFollowList(ctx context.Context, req *follow.FollowListRequest) ([]int64, error) {
	resp, err := followClient.FollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	v := make([]int64, len(resp.UserList))
	for _, user := range resp.UserList {
		v = append(v, user.Id)
	}
	return v, nil
}
