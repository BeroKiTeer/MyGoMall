package service

import (
	"context"
	"fmt"
	"product/biz/dal"
	"product/biz/dal/mysql"
	"product/biz/model"
	"product/kitex_gen/product"
	"strconv"
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
	p := getProduct(int(req.GetId()))
	resp = &product.GetProductResp{}
	resp.Product = &product.Product{
		Id:          uint32(p.Id),
		Name:        p.Name,
		Description: p.Description,
		Picture:     p.Images,
		Price:       p.Price,
		Categories:  []string{},
	}
	return resp, nil
}
func getProduct(id int) model.Product {
	dal.Init()
	var row model.Product
	mysql.DB.Model(&model.Product{}).Where("id=" + strconv.Itoa(id)).Find(&row)
	fmt.Printf("%+v\n", row)
	fmt.Println("----")
	return row
}
