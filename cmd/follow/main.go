package main

import (
	"github.com/benxinm/tiktok/cmd/follow/dal"
	"github.com/benxinm/tiktok/cmd/user/rpc"
	"github.com/benxinm/tiktok/config"
	follow "github.com/benxinm/tiktok/kitex_gen/follow/followservice"
	"github.com/benxinm/tiktok/pkg/constants"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func Init() {
	config.InitLocal("../deploy/config/config.yml", constants.FollowServiceName)
	dal.Init()
	rpc.Init()
}

func main() {
	Init()
	_, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})
	if err != nil {
		panic(err)
	}
	svr := follow.NewServer(new(FollowServiceImpl))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
