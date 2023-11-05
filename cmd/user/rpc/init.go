package rpc

import (
	"github.com/benxinm/tiktok/kitex_gen/follow/followservice"
	"github.com/benxinm/tiktok/kitex_gen/video/videoservice"
)

var (
	videoClient  videoservice.Client
	followClient followservice.Client
)

func Init() {
	InitFollowRpc()
	InitVideoRpc()
}
