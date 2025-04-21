package protocol

import (
	"context"
	"fmt"

	"kitex-multi-protocol/kitex_gen/user"
)

// ThriftHandlerImpl implements the ThriftHandler interface.
type ThriftHandlerImpl struct {
	Service user.UserService
}

// HandleThriftRequest processes incoming Thrift requests.
func (h *ThriftHandlerImpl) HandleThriftRequest(ctx context.Context, methodName string, args interface{}) (interface{}, error) {
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
