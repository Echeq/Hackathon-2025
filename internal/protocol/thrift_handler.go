package protocol

import (
	"context"
	"log"
	"net"

	"kitex-multi-protocol/kitex_gen/user"
)

// ThriftHandler maneja solicitudes Thrift
type ThriftHandler struct {
	Service user.UserService
}

func (h *ThriftHandler) Handle(ctx context.Context, conn net.Conn) error {
	log.Println("Handling Thrift request...")
	return nil
}
