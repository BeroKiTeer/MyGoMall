package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/auth"
	"github.com/cloudwego/kitex/pkg/klog"
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
		klog.Error("获取token失败", err)
		return nil, err
	}

	// 续期，其实就是重新生成一个 token
	newToken, err := GenerateJWT(userID, req.Seconds, s.ctx)
	if err != nil {
		klog.Error("续期token失败", err)
		return nil, err
	}

	return &auth.RefreshTokenResp{NewToken: newToken}, err
}
