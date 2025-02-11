package service

import (
	"context"
	"fmt"
	"product/biz/dal"
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
	p, categories := getProduct(int(req.GetId()))
	resp = &product.GetProductResp{}
	resp.Product = &product.Product{
		Id:          uint32(p.ID),
		Name:        p.Name,
		Description: p.Description,
		Picture:     p.Images,
		Price:       p.Price,
		Categories:  categories,
	}
	return resp, nil
}
func getProduct(id int) (model.Product, []string) {
	dal.Init()
	var row model.Product
	db := mysql.DB
	db.Model(&model.Product{}).Where("id=?", id).Find(&row)
	var categoriesProduct []int
	db.Table("category_product").
		Select("category_id").
		Where("product_id= ?", id).Find(&categoriesProduct)
	var categories []string
	db.Table("categories").
		Select("name").
		Where("id in ?", categoriesProduct).
		Find(&categories)
	fmt.Printf("%+v\n", row)
	fmt.Println("----")
	return row, categories
}
