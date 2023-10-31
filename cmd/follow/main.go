package main

import (
	follow "github.com/benxinm/tiktok/kitex_gen/follow/followservice"
	"log"
)

func main() {
	svr := follow.NewServer(new(FollowServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
