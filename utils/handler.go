package main

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net"
    "net/http"
    "regexp"

    "kitex-multi-protocol/kitex_gen/user"
)

// Handler maneja las solicitudes HTTP y Thrift.
type Handler struct {
    Service user.UserService
}

// NewHandler crea una nueva instancia de Handler.
func NewHandler(service user.UserService) *Handler {
    return &Handler{
        Service: service,
    }
}

// DetectProtocol detecta si una solicitud es Thrift o HTTP.
func DetectProtocol(conn net.Conn) (string, error) {
    buffer := make([]byte, 4)
    _, err := conn.Read(buffer)
    if err != nil {
        return "", err
    }

    conn = &bufferedConn{Conn: conn, buffer: buffer}

    httpReg := regexp.MustCompile(`^(?:GET|POST|PUT|DELETE|HEAD|OPTIONS|CONNECT|TRACE|PATCH)`)
    if httpReg.Match(buffer) {
        return "HTTP", nil
    }
    return "Thrift", nil
}

// HandleHTTPRequest maneja las solicitudes HTTP.
func (h *Handler) HandleHTTPRequest(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/api/UserService/GetUser" && r.Method == "POST" {
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Error al leer el cuerpo", http.StatusBadRequest)
            return
        }

        var params map[string]interface{}
        err = json.Unmarshal(body, &params)
        if err != nil {
            http.Error(w, "Error al parsear JSON", http.StatusBadRequest)
            return
        }

        userID, ok := params["userID"].(float64)
        if !ok {
            http.Error(w, "userID no válido", http.StatusBadRequest)
            return
        }

        result, err := h.Service.GetUser(context.Background(), int64(userID))
        if err != nil {
            http.Error(w, "Error al procesar la solicitud", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"result": result})
    } else {
        http.Error(w, "Ruta no encontrada", http.StatusNotFound)
    }
}

// HandleThriftRequest maneja las solicitudes Thrift.
func (h *Handler) HandleThriftRequest(ctx context.Context, methodName string, args interface{}) (interface{}, error) {
    switch methodName {
    case "GetUser":
        userID, ok := args.(int64)
        if !ok {
            return nil, fmt.Errorf("argumento inválido para GetUser")
        }
        return h.Service.GetUser(ctx, userID)
    default:
        return nil, fmt.Errorf("método no encontrado: %s", methodName)
    }
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