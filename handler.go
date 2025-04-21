package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/remote"
)

// DualProtocolHandler supports HTTP and Thrift traffic on one port
type DualProtocolHandler struct {
	delegate    remote.ServerTransHandler
	httpHandler http.Handler
}

func NewDualProtocolHandler(delegate remote.TransHandler, httpHandler http.Handler) *DualProtocolHandler {
	return &DualProtocolHandler{
		delegate:    delegate,
		httpHandler: httpHandler,
	}
}

func (h *DualProtocolHandler) OnRead(ctx context.Context, conn net.Conn) error {
	buf := make([]byte, 4)
	_, err := conn.Read(buf)
	if err != nil {
		return err
	}
	conn = &peekedConn{Conn: conn, peeked: buf}

	if isHTTPRequest(buf) {
		return h.handleHTTP(conn)
	}
	return h.delegate.OnRead(ctx, conn)
}

func (h *DualProtocolHandler) OnActive(ctx context.Context, conn net.Conn) error {
	return h.delegate.OnActive(ctx, conn)
}

func (h *DualProtocolHandler) OnInactive(ctx context.Context, conn net.Conn) error {
	return h.delegate.OnInactive(ctx, conn)
}

func (h *DualProtocolHandler) OnError(ctx context.Context, err error, conn net.Conn) {
	h.delegate.OnError(ctx, err, conn)
}

func isHTTPRequest(buf []byte) bool {
	return bytes.HasPrefix(buf, []byte("GET ")) ||
		bytes.HasPrefix(buf, []byte("POST")) ||
		bytes.HasPrefix(buf, []byte("PUT ")) ||
		bytes.HasPrefix(buf, []byte("HTTP"))
}

func (h *DualProtocolHandler) handleHTTP(conn net.Conn) error {
	req, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		return err
	}
	defer req.Body.Close()

	response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/json\r\n\r\n%s",
		`{"message":"hello from server"}`)

	_, err = conn.Write([]byte(response))
	return err
}

type peekedConn struct {
	net.Conn
	peeked []byte
}

func (c *peekedConn) Read(p []byte) (n int, err error) {
	if len(c.peeked) > 0 {
		n = copy(p, c.peeked)
		c.peeked = c.peeked[n:]
		return n, nil
	}
	return c.Conn.Read(p)
}
