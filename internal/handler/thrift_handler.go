package handler

import (
    "context"
    "fmt"

    "kitex-multi-protocol/kitex_gen/user"
)

// ThriftHandler maneja las solicitudes Thrift.
type ThriftHandler struct {
    Service user.UserService
}

// NewThriftHandler crea una nueva instancia de ThriftHandler.
func NewThriftHandler(service user.UserService) *ThriftHandler {
    return &ThriftHandler{
        Service: service,
    }
}

// HandleThriftRequest procesa una solicitud Thrift.
func (h *ThriftHandler) HandleThriftRequest(ctx context.Context, methodName string, args interface{}) (interface{}, error) {
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