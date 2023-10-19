package main

import (
	user "github.com/benxinm/tiktok/kitex_gen/user/userservice"
	"github.com/benxinm/tiktok/pkg/constants"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"log"
)

func main() {
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
