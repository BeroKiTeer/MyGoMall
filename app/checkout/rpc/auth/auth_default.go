package auth

import (
	"context"
	auth "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/auth"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq, callOptions ...callopt.Option) (resp *auth.DeliveryResp, err error) {
	resp, err = defaultClient.DeliverTokenByRPC(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeliverTokenByRPC call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq, callOptions ...callopt.Option) (resp *auth.VerifyResp, err error) {
	resp, err = defaultClient.VerifyTokenByRPC(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VerifyTokenByRPC call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func RefreshToken(ctx context.Context, req *auth.RefreshTokenReq, callOptions ...callopt.Option) (resp *auth.RefreshTokenResp, err error) {
	resp, err = defaultClient.RefreshToken(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "RefreshToken call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DecodeToken(ctx context.Context, req *auth.DecodeTokenReq, callOptions ...callopt.Option) (resp *auth.DecodeTokenResp, err error) {
	resp, err = defaultClient.DecodeToken(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DecodeToken call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
