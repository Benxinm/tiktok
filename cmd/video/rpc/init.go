package rpc

import (
	"github.com/benxinm/tiktok/kitex_gen/interaction/interactionservice"
	"github.com/benxinm/tiktok/kitex_gen/user/userservice"
)

var (
	userClient        userservice.Client
	interactionClient interactionservice.Client
)

func Init() {
	InitUserRpc()

}
