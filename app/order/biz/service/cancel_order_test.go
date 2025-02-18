package service

import (
	"context"
	order "order/kitex_gen/order"
	"testing"
)

func TestCancelOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCancelOrderService(ctx)
	// init req and assert value

	req := &order.CancelOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
