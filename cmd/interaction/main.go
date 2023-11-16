package main

import (
	"github.com/benxinm/tiktok/cmd/interaction/dal"
	"github.com/benxinm/tiktok/cmd/interaction/rpc"
	"github.com/benxinm/tiktok/config"
	interaction "github.com/benxinm/tiktok/kitex_gen/interaction/interactionservice"
	"github.com/benxinm/tiktok/pkg/constants"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func Init() {
	config.InitLocal("../deploy/config/config.yml", constants.InteractionServiceName)
	rpc.Init()
	dal.Init()
}
func main() {
	Init()
	_, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})
	if err != nil {
		panic(err)
	}
	svr := interaction.NewServer(new(InteractionServiceImpl))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
