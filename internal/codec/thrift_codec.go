package codec

import (
	"context"
)

type ThriftCodec struct{}

func (t *ThriftCodec) Encode(ctx context.Context, req interface{}) ([]byte, error) {
	// Implementación de codificación Thrift
	return nil, nil
}

func (t *ThriftCodec) Decode(ctx context.Context, data []byte) (interface{}, error) { // Implementación de decodificación Thrift
	return nil, nil
}
