package service

import (
	cart "cart/kitex_gen/cart"
	"context"
	"testing"
)

func TestAddItem_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddItemService(ctx)
	// init req and assert value

	req := &cart.AddItemReq{UserId: 123, Item: &cart.CartItem{ProductId: 456, Quantity: 7}}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
