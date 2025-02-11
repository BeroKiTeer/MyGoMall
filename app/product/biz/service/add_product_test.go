package service

import (
	"context"
	"product/biz/dal"
	product "product/kitex_gen/product"
	"testing"
)

// 测试添加商品代码
func TestAddProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddProductService(ctx)
	dal.Init()
	// init req and assert value
	req := &product.AddProductReq{
		CategoryId:    int32(1),
		Name:          "1",
		Description:   "1",
		Price:         float32(2),
		OriginalPrice: float32(1),
		Stock:         int32(1),
		Images:        "1",
		Status:        int32(1),
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
