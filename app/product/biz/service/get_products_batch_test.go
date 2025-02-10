package service

import (
	"context"
	product "product/kitex_gen/product"
	"testing"
)

func TestGetProductsBatch_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetProductsBatchService(ctx)
	// init req and assert value

	req := &product.GetProductsBatchReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
