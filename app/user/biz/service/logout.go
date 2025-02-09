package service

import (
	"context"
	user "user/kitex_gen/user"
)

type LogoutService struct {
	ctx context.Context
} // NewLogoutService new LogoutService
func NewLogoutService(ctx context.Context) *LogoutService {
	return &LogoutService{ctx: ctx}
}

// Run create note info
func (s *LogoutService) Run(req *user.LogoutReq) (resp *user.LogoutResp, err error) {
	// Finish your business logic.

	return
}
