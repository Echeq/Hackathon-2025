package multiprotocol

import (
    "context"
    "net"
    "strings"

    "github.com/cloudwego/kitex/pkg/transport"
)

type MultiProtocolTransHandlerFactory struct{}

func (f *MultiProtocolTransHandlerFactory) NewTransHandler(opt *transport.ServerTransHandlerOptions) (transport.ServerTransHandler, error) {
    return &MultiProtocolTransHandler{}, nil
}

type MultiProtocolTransHandler struct{}

func (h *MultiProtocolTransHandler) OnActive(ctx context.Context, conn net.Conn) (context.Context, error) {
    buffer := make([]byte, 4)
    _, err := conn.Read(buffer)
    if err != nil {
        return ctx, err
    }

    if strings.HasPrefix(string(buffer), "POST") || strings.HasPrefix(string(buffer), "GET") {
        ctx = context.WithValue(ctx, "protocol", "http")
    } else {
        ctx = context.WithValue(ctx, "protocol", "thrift")
    }
    return ctx, nil
}