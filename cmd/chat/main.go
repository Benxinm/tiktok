package main

import (
	"github.com/benxinm/tiktok/cmd/chat/dal"
	"github.com/benxinm/tiktok/config"
	chat "github.com/benxinm/tiktok/kitex_gen/chat/messageservice"
	"github.com/benxinm/tiktok/pkg/constants"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func Init() {
	config.InitLocal("../deploy/config/config.yml", constants.ChatServiceName)
	dal.Init()
}
func main() {
	Init()
	_, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr})
	if err != nil {
		panic(err)
	}

	svr := chat.NewServer(new(MessageServiceImpl))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
