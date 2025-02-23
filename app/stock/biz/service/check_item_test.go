package service

import (
	"context"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	"stock/biz/dal/mysql"
	"testing"
)

func TestCheckItem_Run(t *testing.T) {
	mysql.Init()
	ctx := context.Background()
	s := NewCheckItemService(ctx)
	// init req and assert value

	req := &stock.CheckItemReq{ProductId: 1}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
