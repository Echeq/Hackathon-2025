package protocol

import "net"

// BufferedConn wraps a net.Conn and allows restoring unread bytes.
type BufferedConn struct {
	net.Conn
	Buffer []byte
}

func NewBufferedConn(conn net.Conn) *BufferedConn {
	return &BufferedConn{
		Conn:   conn,
		Buffer: make([]byte, 4),
	}
}

func (b *BufferedConn) Read(p []byte) (int, error) {
	if len(b.Buffer) > 0 {
		// Leer primero desde el buffer
		n := copy(p, b.Buffer)
		b.Buffer = b.Buffer[n:]
		return n, nil
	}
	// Si no hay más datos en el buffer, leer directamente de la conexión
	return b.Conn.Read(p)
}
