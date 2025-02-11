package service

import (
	auth "auth/kitex_gen/auth"
	"context"
)

type GetTokenService struct {
	ctx context.Context
} // NewGetTokenService new GetTokenService
func NewGetTokenService(ctx context.Context) *GetTokenService {
	return &GetTokenService{ctx: ctx}
}

// Run create note info
func (s *GetTokenService) Run(req *auth.GetTokenReq) (resp *auth.GetTokenResp, err error) {
	// Finish your business logic.

	return
}
