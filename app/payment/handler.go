package main

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/payment"
	"payment/biz/service"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	resp, err = service.NewChargeService(ctx).Run(req)

	return resp, err
}

// CancelPayment implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) CancelPayment(ctx context.Context, req *payment.CancelReq) (resp *payment.CancelResp, err error) {
	resp, err = service.NewCancelPaymentService(ctx).Run(req)

	return resp, err
}
