package main

import (
	"context"
	"log"
	"net"

	"kitex-multi-protocol/internal/protocol"
	"kitex-multi-protocol/kitex_gen/user"
	"kitex-multi-protocol/utils"
)

func main() {
	// Crear una instancia del servicio
	service := utils.NewUserServiceImpl()

	// Crear una fábrica de manejadores
	transHandlerFactory := protocol.NewTransHandlerFactory()

	// Configurar el servidor
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		go handleConnection(conn, transHandlerFactory, service)
	}
}

func handleConnection(conn net.Conn, factory *protocol.TransHandlerFactory, service user.UserService) {
	defer conn.Close()

	// Detectar el protocolo usando TransHandlerFactory
	pre := make([]byte, 4)
	n, err := conn.Read(pre)
	if err != nil {
		log.Printf("Error detecting protocol: %v", err)
		return
	}
	if n < 4 {
		log.Printf("Error detecting protocol: not enough data to detect protocol")
		return
	}

	protocolType, err := factory.ProtocolMatchFromPreRead(pre)
	if err != nil {
		log.Printf("Error detecting protocol: %v", err)
		return
	}

	// Restaurar los bytes leídos durante la detección del protocolo
	bufferedConn := protocol.NewBufferedConn(conn)
	bufferedConn.Buffer = pre
	if n > 0 {
		bufferedConn.Buffer = bufferedConn.Buffer[:n]
	}

	// Crear el handler adecuado
	handler := protocol.CreateHandler(protocolType, service)
	if handler == nil {
		log.Printf("No handler found for protocol: %s", protocolType)
		return
	}

	// Manejar la solicitud usando la conexión envuelta
	ctx := context.Background()
	if err := handler.Handle(ctx, bufferedConn); err != nil {
		log.Printf("Error handling request: %v", err)
	}
}
