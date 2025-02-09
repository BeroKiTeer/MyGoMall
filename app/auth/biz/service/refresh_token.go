package service

import (
	auth "auth/kitex_gen/auth"
	"context"
)

type RefreshTokenService struct {
	ctx context.Context
} // NewRefreshTokenService new RefreshTokenService
func NewRefreshTokenService(ctx context.Context) *RefreshTokenService {
	return &RefreshTokenService{ctx: ctx}
}

// Run create note info
func (s *RefreshTokenService) Run(req *auth.RefreshTokenReq) (resp *auth.RefreshTokenResp, err error) {
	// Finish your business logic.

	return
}
