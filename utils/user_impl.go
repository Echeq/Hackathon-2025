package main

import (
    "context"
    "fmt"

    "kitex-multi-protocol/kitex_gen/user"
)

// UserServiceImpl implementa la interfaz UserService.
type UserServiceImpl struct{}

// GetUser implementa el método GetUser de la interfaz UserService.
func (s *UserServiceImpl) GetUser(ctx context.Context, userID int64) (string, error) {
    return fmt.Sprintf("Usuario con ID %d", userID), nil
}