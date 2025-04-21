package main

import (
    "bufio"
    "log"
    "net"
    "net/http"

    "kitex-multi-protocol/utils" // Importamos el paquete utils
)

func main() {
    // Create an instance of the service
    service := utils.NewUserServiceImpl()

    // Create a handler
    handler := utils.NewHandler(service)

    // Configure the server
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

        go handleConnection(conn, handler)
    }
}

func handleConnection(conn net.Conn, handler *utils.Handler) {
    protocolType, err := utils.DetectProtocol(conn)
    if err != nil {
        log.Printf("Error detecting protocol: %v", err)
        conn.Close()
        return
    }

    if protocolType == "HTTP" {
        log.Println("HTTP request detected")
        request, err := http.ReadRequest(bufio.NewReader(conn))
        if err != nil {
            log.Printf("Error reading HTTP request: %v", err)
            conn.Close()
            return
        }

        response := http.Response{
            Status:        "200 OK",
            StatusCode:    200,
            Proto:         "HTTP/1.1",
            ProtoMajor:    1,
            ProtoMinor:    1,
            Body:          nil,
            ContentLength: 0,
        }

        handler.HandleHTTPRequest(&responseWriter{Conn: conn}, request)

        // Write the HTTP response to the connection
        response.Write(conn)
        conn.Close()
    } else if protocolType == "Thrift" {
        log.Println("Thrift request detected")
        // Here you would handle Thrift requests using handler.HandleThriftRequest
    }
}

type responseWriter struct {
    net.Conn
    header http.Header
}

func (w *responseWriter) Header() http.Header {
    if w.header == nil {
        w.header = make(http.Header)
    }
    return w.header
}

func (w *responseWriter) Write(data []byte) (int, error) {
    return w.Conn.Write(data)
}

func (w *responseWriter) WriteHeader(statusCode int) {
    // No need to do anything here since we are writing directly to the connection.
}