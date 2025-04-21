package codec

import (
	"context"
	"encoding/json"
)

type HTTPCodec struct{}

func (h *HTTPCodec) Encode(ctx context.Context, req interface{}) ([]byte, error) { // Implementación de codificación HTTP
	return nil, nil
}

func (h *HTTPCodec) Decode(ctx context.Context, data []byte) (interface{}, error) {
	var params map[string]interface{}
	err := json.Unmarshal(data, &params)
	if err != nil {
		return nil, err
	}
	return params, nil
}
