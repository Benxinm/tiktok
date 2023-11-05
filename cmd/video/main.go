package main

import (
	"github.com/benxinm/tiktok/cmd/video/dal"
	"github.com/benxinm/tiktok/cmd/video/rpc"
	video "github.com/benxinm/tiktok/kitex_gen/video/videoservice"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func Init() {
	dal.Init()
	rpc.Init()
}

func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:80"}) //TODO config file
	if err != nil {
		panic(err)
	}

	svr := video.NewServer(new(VideoServiceImpl))

	if err := svr.Run(); err != nil {
		log.Println(err.Error())
		panic(err)
	}
}
