package main

import (
	interaction "github.com/benxinm/tiktok/kitex_gen/interaction/interactionservice"
	"log"
)

func main() {
	svr := interaction.NewServer(new(InteractionServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
