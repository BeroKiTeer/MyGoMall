package service

import (
	"context"
	"product/biz/dal/mysql"
	"product/kitex_gen/product"
	"testing"
)

func TestCreateProduct_Run(t *testing.T) {
	mysql.Init()
	ctx := context.Background()
	s := NewCreateProductService(ctx)
	//init req and assert value
	newProduct := &product.Product{
		Categories:    []string{"1", "2", "3"},
		Name:          "1",
		Description:   "1",
		Price:         float32(2),
		OriginalPrice: float32(1),
		Stock:         uint32(1),
		Images:        "1",
		Status:        uint32(1),
	}
	req := &product.CreateProductReq{
		Product: newProduct,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
