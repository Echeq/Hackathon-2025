package utils

import (
    "context"
    "fmt"
)

// UserServiceImpl implements the UserService interface.
type UserServiceImpl struct{}

// GetUser implements the GetUser method of the UserService interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, userID int64) (string, error) {
    return fmt.Sprintf("User with ID %d", userID), nil
}

// NewUserServiceImpl creates a new instance of UserServiceImpl.
func NewUserServiceImpl() *UserServiceImpl {
    return &UserServiceImpl{}
}