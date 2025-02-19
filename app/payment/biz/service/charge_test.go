package service

import (
	"context"
	"payment/biz/dal/mysql"
	payment "payment/kitex_gen/payment"
	"testing"
)

func TestCharge_Run(t *testing.T) {
	mysql.Init()
	ctx := context.Background()
	s := NewChargeService(ctx)
	// init req and assert value

	req := &payment.ChargeReq{
		UserId: 1, OrderId: "2", Amount: 1, CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          "4242424242424242",
			CreditCardCvv:             1111,
			CreditCardExpirationYear:  2025,
			CreditCardExpirationMonth: 3,
		},
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
