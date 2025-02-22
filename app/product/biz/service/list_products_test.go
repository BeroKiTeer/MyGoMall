package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/product"
	"testing"
)

func TestListProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListProductsService(ctx)
	// init req and assert value

	req := &product.ListProductsReq{
		Page:         2,
		PageSize:     10,
		CategoryName: "测试",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
