package main

import (
	"auth/biz/service"
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/auth"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	resp, err = service.NewDeliverTokenByRPCService(ctx).Run(req)

	return resp, err
}

// VerifyTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	resp, err = service.NewVerifyTokenByRPCService(ctx).Run(req)

	return resp, err
}

// RefreshToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) RefreshToken(ctx context.Context, req *auth.RefreshTokenReq) (resp *auth.RefreshTokenResp, err error) {
	resp, err = service.NewRefreshTokenService(ctx).Run(req)

	return resp, err
}

// DecodeToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DecodeToken(ctx context.Context, req *auth.DecodeTokenReq) (resp *auth.DecodeTokenResp, err error) {
	resp, err = service.NewDecodeTokenService(ctx).Run(req)

	return resp, err
}
