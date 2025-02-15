package service

import (
	cart "cart/kitex_gen/cart"
	"cart/rpc"
	"context"
	"testing"
)

func TestAddItem_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddItemService(ctx)
	// init req and assert value

	rpc.InitClient()
	req := &cart.AddItemReq{UserId: 121, Item: &cart.CartItem{ProductId: 1, Quantity: 1}}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
