package utils

import (
	"context"
	"strconv" // Importamos el paquete strconv para manejar la conversión
)

// UserServiceImpl implements the UserService interface.
type UserServiceImpl struct{}

// NewUserServiceImpl creates a new instance of UserServiceImpl.
func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{}
}

// GetUser returns a user based on the provided userID.
func (s *UserServiceImpl) GetUser(ctx context.Context, userID int64) (string, error) {
	// Convertimos el userID a una cadena usando strconv.FormatInt
	return "Usuario con ID " + strconv.FormatInt(userID, 10), nil
}
