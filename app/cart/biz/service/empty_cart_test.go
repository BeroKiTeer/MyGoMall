package service

import (
	"cart/biz/dal/mysql"
	cart "cart/kitex_gen/cart"
	"context"
	"testing"
)

func TestEmptyCart_Run(t *testing.T) {
	mysql.Init()
	ctx := context.Background()
	s := NewEmptyCartService(ctx)
	// init req and assert value

	req := &cart.EmptyCartReq{UserId: 121}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
