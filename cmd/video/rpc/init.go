package rpc

import (
	"github.com/benxinm/tiktok/kitex_gen/follow/followservice"
	"github.com/benxinm/tiktok/kitex_gen/interaction/interactionservice"
	"github.com/benxinm/tiktok/kitex_gen/user/userservice"
)

var (
	userClient        userservice.Client
	interactionClient interactionservice.Client
	followClient      followservice.Client
)

func Init() {
	InitUserRpc()
	InitInteractionRpc()
	InitFollowRpc()
}
