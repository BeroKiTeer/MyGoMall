package service

import (
	"context"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	"testing"
)

func TestReserveItem_Run(t *testing.T) {
	ctx := context.Background()
	s := NewReserveItemService(ctx)
	// init req and assert value

	req := &stock.ReserveItemReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
