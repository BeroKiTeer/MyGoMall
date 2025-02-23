package service

import (
	"context"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	"testing"
)

func TestReduceItem_Run(t *testing.T) {
	ctx := context.Background()
	s := NewReduceItemService(ctx)
	// init req and assert value

	req := &stock.ReduceItemReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
