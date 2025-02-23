package service

import (
	"context"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	"stock/biz/dal/mysql"
	"testing"
)

func TestReduceItem_Run(t *testing.T) {
	mysql.Init()
	ctx := context.Background()
	s := NewReduceItemService(ctx)
	// init req and assert value

	req := &stock.ReduceItemReq{ProductId: 1, Quantity: 1}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
