package utils

import (
    "context"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net"
    "net/http"
    "regexp"

    "kitex-multi-protocol/kitex_gen/user"
)

// Handler handles both HTTP and Thrift requests.
type Handler struct {
    Service user.UserService
}

// NewHandler creates a new instance of Handler.
func NewHandler(service user.UserService) *Handler {
    return &Handler{
        Service: service,
    }
}

// DetectProtocol detects whether a request is HTTP or Thrift.
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

// HandleHTTPRequest handles HTTP requests.
func (h *Handler) HandleHTTPRequest(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/api/UserService/GetUser" && r.Method == "POST" {
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Error reading request body", http.StatusBadRequest)
            return
        }

        var params map[string]interface{}
        err = json.Unmarshal(body, &params)
        if err != nil {
            http.Error(w, "Error parsing JSON", http.StatusBadRequest)
            return
        }

        userID, ok := params["userID"].(float64)
        if !ok {
            http.Error(w, "Invalid userID", http.StatusBadRequest)
            return
        }

        result, err := h.Service.GetUser(context.Background(), int64(userID))
        if err != nil {
            http.Error(w, "Error processing request", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"result": result})
    } else {
        http.Error(w, "Route not found", http.StatusNotFound)
    }
}

// HandleThriftRequest handles Thrift requests.
func (h *Handler) HandleThriftRequest(ctx context.Context, methodName string, args interface{}) (interface{}, error) {
    switch methodName {
    case "GetUser":
        userID, ok := args.(int64)
        if !ok {
            return nil, fmt.Errorf("invalid argument for GetUser")
        }
        return h.Service.GetUser(ctx, userID)
    default:
        return nil, fmt.Errorf("method not found: %s", methodName)
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