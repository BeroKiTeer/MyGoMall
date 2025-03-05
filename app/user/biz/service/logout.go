package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/auth"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"user/rpc"
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
	//先根据用户id获取用户token(覆盖之前的)
	TokenReq := auth.DeliverTokenReq{
		UserId: req.UserId,
	}

	token, err := rpc.AuthClient.DeliverTokenByRPC(s.ctx, &TokenReq)
	if err != nil {
		klog.Errorf("获取token失败：%v", err)
		return nil, err
	}
	//删除token（时间置零）
	refreshReq := auth.RefreshTokenReq{
		Token:   token.Token,
		Seconds: 0,
	}

	_, err = rpc.AuthClient.RefreshToken(s.ctx, &refreshReq)
	if err != nil {
		klog.Error("删除token失败", err)
		return nil, err
	}

	respTemp := user.LogoutResp{
		Success: true,
	}

	return &respTemp, nil

}
