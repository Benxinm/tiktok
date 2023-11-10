package rpc

import (
	"github.com/benxinm/tiktok/kitex_gen/user/userservice"
	"github.com/benxinm/tiktok/kitex_gen/video/videoservice"
)

var (
	videoClient videoservice.Client
	userClient  userservice.Client
)

func Init() {
	InitUserRpc()
	InitVideoRpc()
}
