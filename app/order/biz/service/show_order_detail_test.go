package service

import (
	"context"
	"testing"
	order "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
)

func TestShowOrderDetail_Run(t *testing.T) {
	ctx := context.Background()
	s := NewShowOrderDetailService(ctx)
	// init req and assert value

	req := &order.ShowOrderDetailReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
