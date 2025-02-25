package service

import (
	"context"
	payment "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/payment"
	"payment/biz/dal/mysql"
	"testing"
)

func TestChargeByThirdParty_Run(t *testing.T) {
	mysql.Init()
	ctx := context.Background()
	s := NewChargeByThirdPartyService(ctx)
	// init req and assert value

	req := &payment.ChargeByThirdPartyReq{Amount: 123, UserId: 233430303, OrderId: "1", Way: "YuanShen"}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
