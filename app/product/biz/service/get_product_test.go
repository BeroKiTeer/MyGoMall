package service

import (
	"context"
	product "product/kitex_gen/product"
	"testing"
)

func TestGetProduct_Run(t *testing.T) {
	//mysql.Init()
	ctx := context.Background()
	s := NewGetProductService(ctx)
	// init req and assert value
	req := &product.GetProductReq{}
	req.Id = 1
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
