package service

import (
	"context"
	payment "payment/kitex_gen/payment"
	"testing"
)

func TestChargeByThirdParty_Run(t *testing.T) {
	ctx := context.Background()
	s := NewChargeByThirdPartyService(ctx)
	// init req and assert value

	req := &payment.ChargeByThirdPartyReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
