package main

import (
    "log"

    "github.com/cloudwego/kitex/server"
    "kitex-multi-protocol/kitex_gen/UserService/userservice"
    "kitex-multi-protocol/multiprotocol"
)

func main() {
    // Define el servicio
    svr := userservice.NewServer(
        new(UserServiceImpl),
        server.WithTransHandlerFactory(multiprotocol.MultiProtocolTransHandlerFactory{}),
    )

    // Inicia el servidor
    log.Println("Starting server on port 8888...")
    if err := svr.Run(); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}