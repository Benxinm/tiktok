package rpc

import "github.com/benxinm/tiktok/kitex_gen/user/userservice"

var (
	userClient userservice.Client
)

func Init() {
	InitUserRpc()
}
