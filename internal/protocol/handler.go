package protocol

import (
	"context"
	"net"
)

// Handler es la interfaz para manejar solicitudes
type Handler interface {
	Handle(ctx context.Context, conn net.Conn) error
}
