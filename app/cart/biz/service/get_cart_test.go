package service

import (
	"cart/biz/dal"
	"cart/biz/dal/mysql"
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/cart"
	"testing"
)

func TestGetCart_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewGetCartService(ctx)
	// init req and assert value

	mysql.Init()
	req := &cart.GetCartReq{UserId: 123}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
