package service

import (
	"cart/biz/dal"
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/cart"
	"testing"
)

func TestAddItem_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewAddItemService(ctx)
	// init req and assert value

	req := &cart.AddItemReq{UserId: 123, Item: &cart.CartItem{ProductId: 456, Quantity: 7}}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
