package main

import (
	"context"
	user "kitex-multi-protocol/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, iD string) (resp *user.User, err error) {
	// TODO: Your code here...
	return
}
