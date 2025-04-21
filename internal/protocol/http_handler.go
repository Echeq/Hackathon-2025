package protocol

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"kitex-multi-protocol/kitex_gen/user"
)

// HTTPHandler maneja solicitudes HTTP
type HTTPHandler struct {
	Service user.UserService
}

func (h *HTTPHandler) Handle(ctx context.Context, conn net.Conn) error {
	httpHandler := &HTTPHandlerImpl{Service: h.Service}
	srv := &http.Server{
		Handler: httpHandler,
	}

	tempListener := newSingleConnListener(conn)
	defer tempListener.Close()

	errChan := make(chan error, 1)
	go func() {
		errChan <- srv.Serve(tempListener)
	}()

	select {
	case <-ctx.Done():
		srv.Shutdown(context.Background())
		return ctx.Err()
	case err := <-errChan:
		if err != nil && err != http.ErrServerClosed {
			log.Printf("Error serving HTTP request: %v", err)
		}
	case <-time.After(5 * time.Second):
		srv.Shutdown(context.Background())
	}

	return nil
}

// singleConnListener es una implementación de net.Listener que solo acepta una conexión
type singleConnListener struct {
	conn net.Conn
}

func newSingleConnListener(conn net.Conn) *singleConnListener {
	return &singleConnListener{conn: conn}
}

func (l *singleConnListener) Accept() (net.Conn, error) {
	if l.conn == nil {
		return nil, net.ErrClosed
	}
	conn := l.conn
	l.conn = nil
	return conn, nil
}

func (l *singleConnListener) Close() error {
	if l.conn != nil {
		err := l.conn.Close()
		l.conn = nil
		return err
	}
	return nil
}

func (l *singleConnListener) Addr() net.Addr {
	if l.conn != nil {
		return l.conn.LocalAddr()
	}
	return nil
}

// HTTPHandlerImpl es la implementación real para manejar solicitudes HTTP
type HTTPHandlerImpl struct {
	Service user.UserService
}

func (h *HTTPHandlerImpl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Aquí debes implementar la lógica para manejar la solicitud HTTP
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("HTTP request handled"))
}
