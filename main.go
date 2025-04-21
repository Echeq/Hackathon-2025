package main

import (
	"log"
	"net/http"

	example "kitex-multi-protocol/kitex_gen/hello/example/helloservice"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/server"
)

func main() {
	// Construct internal thrift handler
	thriftHandler := remote.NewTransServerHandler(remote.ServerOption{})

	// Construct dual handler with HTTP and thrift support
	dualHandler := NewDualProtocolHandler(thriftHandler, http.NewServeMux()) // httpHandler can be customized

	// Create the Kitex server and pass the custom transport handler
	svr := example.NewServer(
		new(HelloServiceImpl),
		server.WithTransHandler(dualHandler),
	)

	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}
