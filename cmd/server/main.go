package main

import (
    "bufio"
    "log"
    "net"
    "net/http"

	"kitex-multi-protocol/utils"
)

func main() {
    // Crear una instancia del servicio
    service := new(UserServiceImpl)

    // Crear un manejador
    handler := NewHandler(service)

    // Configurar el servidor
    listener, err := net.Listen("tcp", ":8888")
    if err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Error al aceptar conexión: %v", err)
            continue
        }

        go handleConnection(conn, handler)
    }
}

func handleConnection(conn net.Conn, handler *Handler) {
    protocolType, err := DetectProtocol(conn)
    if err != nil {
        log.Printf("Error al detectar protocolo: %v", err)
        conn.Close()
        return
    }

    if protocolType == "HTTP" {
        log.Println("Solicitud HTTP detectada")
        request, err := http.ReadRequest(bufio.NewReader(conn))
        if err != nil {
            log.Printf("Error al leer solicitud HTTP: %v", err)
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

        // Escribir la respuesta HTTP en la conexión
        response.Write(conn)
        conn.Close()
    } else if protocolType == "Thrift" {
        log.Println("Solicitud Thrift detectada")
        // Aquí manejarías la solicitud Thrift usando handler.HandleThriftRequest
    }
}

type responseWriter struct {
    net.Conn
    header http.Header
}

// Header devuelve las cabeceras de la respuesta.
func (w *responseWriter) Header() http.Header {
    if w.header == nil {
        w.header = make(http.Header)
    }
    return w.header
}

// Write escribe los datos en la conexión.
func (w *responseWriter) Write(data []byte) (int, error) {
    return w.Conn.Write(data)
}

// WriteHeader escribe el código de estado HTTP.
func (w *responseWriter) WriteHeader(statusCode int) {
    // No necesitamos hacer nada aquí porque estamos escribiendo directamente en la conexión.
}