package main

import (
	example "kitex-multi-protocol/kitex_gen/hello/example/helloservice"
	"log"

	"github.com/cloudwego/kitex/server"
)

func main() {
	svr := example.NewServer(
		new(HelloServiceImpl),
		server.WithTransHandlerFactory(&DualProtocolTransHandlerFactory{}),
	)

	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
