package main

import (
    "log"
    "net"

    "github.com/cloudwego/kitex/server"
    "kitex-multi-protocol/kitex_gen/example/exampleservice"
    "kitex-multi-protocol/biz/service"
)

func main() {
    // Configura la dirección del servidor
    addr, err := net.ResolveTCPAddr("tcp", ":8888")
    if err != nil {
        log.Fatalf("Failed to resolve address: %v", err)
    }

    // Crea el servidor
    svr := exampleservice.NewServer(
        new(service.ExampleServiceImpl),
        server.WithServiceAddr(addr),
    )

    // Inicia el servidor
    log.Println("Starting server on port 8888...")
    if err := svr.Run(); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}