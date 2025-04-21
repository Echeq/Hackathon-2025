package protocol

import (
	"regexp"
)

// TransHandlerFactory encapsulates the logic to detect protocols and create handlers.
type TransHandlerFactory struct{}

// NewTransHandlerFactory creates a new instance of TransHandlerFactory.
func NewTransHandlerFactory() *TransHandlerFactory {
	return &TransHandlerFactory{}
}

// ProtocolMatchFromPreRead detects whether a request is HTTP/1.1 or Thrift from pre-read data.
func (f *TransHandlerFactory) ProtocolMatchFromPreRead(pre []byte) (string, error) {
	httpReg := regexp.MustCompile(`^(?:GET|POST|PUT|DELETE|HEAD|OPTIONS|CONNECT|TRACE|PATCH)`)

	// Convertir pre a una cadena para usar con Match
	if httpReg.MatchString(string(pre)) {
		return "HTTP", nil
	}
	return "Thrift", nil
}
