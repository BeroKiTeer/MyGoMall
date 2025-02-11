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

	// 先解码。获得 userID
	userID, err := GetUserIDFromToken(req.Token)
	if err != nil {
		return nil, err
	}

	// 续期，其实就是重新生成一个 token
	newToken, err := GenerateJWT(userID, req.Seconds)
	if err != nil {
		return nil, err
	}

	return &auth.RefreshTokenResp{NewToken: newToken}, err
}
