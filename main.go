package main

import (
	example "kitex-multi-protocol/kitex_gen/hello/example/helloservice"
	"log"
)

func main() {
	svr := example.NewServer(new(HelloServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
