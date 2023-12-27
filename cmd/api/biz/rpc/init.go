package rpc

import (
	"github.com/benxinm/tiktok/kitex_gen/chat/messageservice"
	"github.com/benxinm/tiktok/kitex_gen/follow/followservice"
	"github.com/benxinm/tiktok/kitex_gen/interaction/interactionservice"
	"github.com/benxinm/tiktok/kitex_gen/video/videoservice"

	"github.com/benxinm/tiktok/kitex_gen/user/userservice"
)

var (
	userClient        userservice.Client
	followClient      followservice.Client
	interactionClient interactionservice.Client
	chatClient        messageservice.Client
	videoClient       videoservice.Client
)

func Init() {
	InitUserRpc()
	InitVideoRpc()
	InitFollowRpc()
	InitChatRpc()
	InitInteractionRpc()
}
