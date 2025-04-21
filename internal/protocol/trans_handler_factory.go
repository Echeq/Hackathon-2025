package protocol

import "kitex-multi-protocol/kitex_gen/user"

// CreateHandler crea un manejador basado en el protocolo detectado
func CreateHandler(protocol string, service user.UserService) Handler {
	switch protocol {
	case "HTTP":
		return &HTTPHandler{Service: service}
	case "Thrift":
		return &ThriftHandler{Service: service}
	default:
		return nil
	}
}
