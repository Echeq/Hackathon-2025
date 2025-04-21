package main

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/kitex/pkg/remote/trans/handler" // Use Kitex default thrift handler
)

// This creates your handler — Kitex will call this
type DualProtocolTransHandlerFactory struct{}

// Match the new expected method signature
func (f *DualProtocolTransHandlerFactory) NewTransHandler(opt *remote.ServerOption) (remote.ServerTransHandler, error) {
	// Use Kitex's built-in Thrift handler for fallback
	return &DualProtocolHandler{
		thriftHandler: handler.NewDefaultServerTransHandlerFactory().NewTransHandler(opt),
	}, nil
}

// This is your custom protocol handler
type DualProtocolHandler struct {
	thriftHandler remote.ServerTransHandler // Holds the default Kitex thrift handler
}

// This handles incoming requests and decides HTTP vs Thrift
func (h *DualProtocolHandler) Read(ctx context.Context, conn net.Conn, msg remote.Message) (context.Context, error) {
	reader := bufio.NewReader(conn)

	headerBytes, err := reader.Peek(4)
	if err != nil {
		return ctx, err
	}

	// Detect HTTP
	if bytes.HasPrefix(headerBytes, []byte("POST")) || bytes.HasPrefix(headerBytes, []byte("GET ")) {
		fmt.Println("Detected: HTTP")
		return ctx, errors.New("HTTP not supported yet") // you'll handle HTTP later
	}

	// Detect Thrift binary
	if headerBytes[0] == 0x80 || headerBytes[0] == 0x82 {
		fmt.Println("Detected: Thrift")
		return h.thriftHandler.Read(ctx, conn, msg)
	}

	return ctx, errors.New("Unknown protocol")
}

// This handles sending response (just let thrift do it)
func (h *DualProtocolHandler) Write(ctx context.Context, conn net.Conn, msg *remote.Message) error {
	return h.thriftHandler.Write(ctx, conn, msg)
}
