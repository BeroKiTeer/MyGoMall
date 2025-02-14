package service

import (
	"context"
	"fmt"
	"product/biz/dal/mysql"
	"product/biz/model"
	"product/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	fmt.Printf("请求id：%+v\n", req.GetId())
	var categories []string

	p, categories, err := model.GetProductWithCategory(mysql.DB, int(req.GetId()))
	if err != nil {
		return nil, err
	}
	resp = &product.GetProductResp{}
	resp.Product = &product.Product{
		Id:            uint32(p.ID),
		Name:          p.Name,
		Description:   p.Description,
		Images:        p.Images,
		Price:         p.Price,
		Categories:    categories,
		OriginalPrice: p.OriginalPrice,
		Stock:         p.Stock,
		Status:        uint32(p.Status),
	}
	return resp, nil
}
