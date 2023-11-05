package main

import (
	"github.com/benxinm/tiktok/cmd/user/rpc"
	follow "github.com/benxinm/tiktok/kitex_gen/follow/followservice"
	"log"
)

func Init() {
	rpc.Init()
}

func main() {
	Init()
	svr := follow.NewServer(new(FollowServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
