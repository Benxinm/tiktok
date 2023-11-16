package main

import (
	"github.com/benxinm/tiktok/cmd/video/dal"
	"github.com/benxinm/tiktok/cmd/video/rpc"
	"github.com/benxinm/tiktok/config"
	video "github.com/benxinm/tiktok/kitex_gen/video/videoservice"
	"github.com/benxinm/tiktok/pkg/constants"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func Init() {
	config.InitLocal("../deploy/config/config.yml", constants.VideoServiceName)
	dal.Init()
	rpc.Init()
}

func main() {
	Init()
	_, err := etcd.NewEtcdRegistry([]string{config.Etcd.Addr}) //TODO config file
	if err != nil {
		panic(err)
	}

	svr := video.NewServer(new(VideoServiceImpl))

	if err := svr.Run(); err != nil {
		log.Println(err.Error())
		panic(err)
	}
}
