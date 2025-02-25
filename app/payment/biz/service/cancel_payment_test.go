package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/payment"
	"payment/biz/dal/mysql"
	"testing"
)

func TestCancelPayment_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCancelPaymentService(ctx)
	// init req and assert value
	mysql.Init()
	req := &payment.CancelReq{
		Id: 1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
