package protocol

import (
	"context"
	"log"
	"net"

	"kitex-multi-protocol/kitex_gen/user"
)

// Handler defines the interface for handling requests.
type Handler interface {
	Handle(ctx context.Context, conn net.Conn) error
}

// HTTPHandler implements the Handler interface for HTTP requests.
type HTTPHandler struct {
	Service user.UserService
}

func (h *HTTPHandler) Handle(ctx context.Context, conn net.Conn) error {
	// Simulación de manejo de solicitudes HTTP
	log.Println("Handling HTTP request...")
	return nil
}

// ThriftHandler implements the Handler interface for Thrift requests.
type ThriftHandler struct {
	Service user.UserService
}

func (h *ThriftHandler) Handle(ctx context.Context, conn net.Conn) error { // Simulación de manejo de solicitudes Thrift
	log.Println("Handling Thrift request...")
	return nil
}

// CreateHandler creates a handler based on the detected protocol.
func CreateHandler(protocol string, service user.UserService) Handler {
	switch protocol {
	case "HTTP":
		return &HTTPHandler{Service: service}
	case "Thrift":
		return &ThriftHandler{Service: service}
	default:
		return nil
	}
}
