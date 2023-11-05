package main

import (
	"github.com/benxinm/tiktok/cmd/user/dal"
	"github.com/benxinm/tiktok/cmd/user/rpc"
	user "github.com/benxinm/tiktok/kitex_gen/user/userservice"
	"github.com/benxinm/tiktok/pkg/constants"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"log"
)

func Init() {
	rpc.Init()
	dal.Init()
}

func main() {
	Init()
	//addr ,err := net.ResolveTCPAddr("tcp",)
	svr := user.NewServer(
		new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: constants.UserServiceName,
		}),
		server.WithMuxTransport(),
		server.WithLimit(&limit.Option{
			MaxConnections: constants.MaxConnections,
			MaxQPS:         constants.MaxQPS,
		}))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
