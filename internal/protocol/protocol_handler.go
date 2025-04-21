package protocol

import (
    "net"
    "regexp"
)

var httpReg = regexp.MustCompile(`^(?:GET|POST|PUT|DELETE|HEAD|OPTIONS|CONNECT|TRACE|PATCH)`)

// DetectProtocol detecta si una solicitud es Thrift o HTTP.
func DetectProtocol(conn net.Conn) (string, error) {
    buffer := make([]byte, 4)
    _, err := conn.Read(buffer)
    if err != nil {
        return "", err
    }

    // Restaurar los bytes leídos al búfer de la conexión
    conn = &bufferedConn{Conn: conn, buffer: buffer}

    if httpReg.Match(buffer) {
        return "HTTP", nil
    }
    return "Thrift", nil
}

type bufferedConn struct {
    net.Conn
    buffer []byte
}

func (b *bufferedConn) Read(p []byte) (int, error) {
    if len(b.buffer) > 0 {
        n := copy(p, b.buffer)
        b.buffer = b.buffer[n:]
        return n, nil
    }
    return b.Conn.Read(p)
}