package payment

import (
	"context"
	payment "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/payment"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Charge(ctx context.Context, req *payment.ChargeReq, callOptions ...callopt.Option) (resp *payment.ChargeResp, err error) {
	resp, err = defaultClient.Charge(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Charge call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CancelPayment(ctx context.Context, req *payment.CancelReq, callOptions ...callopt.Option) (resp *payment.CancelResp, err error) {
	resp, err = defaultClient.CancelPayment(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CancelPayment call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ChargeByThirdParty(ctx context.Context, req *payment.ChargeByThirdPartyReq, callOptions ...callopt.Option) (resp *payment.ChargeByThirdPartyResp, err error) {
	resp, err = defaultClient.ChargeByThirdParty(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ChargeByThirdParty call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
