package service

import (
	"context"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	"testing"
)

func TestCheckItem_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCheckItemService(ctx)
	// init req and assert value

	req := &stock.CheckItemReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
