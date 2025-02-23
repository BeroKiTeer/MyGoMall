package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/payment"
	"payment/biz/dal/mysql"
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
