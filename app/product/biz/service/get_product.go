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
	p, categories, err := getProduct(int(req.GetId()))
	if err != nil {
		return nil, err
	}
	resp = &product.GetProductResp{}
	resp.Product = &product.Product{
		Id:          uint32(p.ID),
		Name:        p.Name,
		Description: p.Description,
		Images:      p.Images,
		Price:       p.Price,
		Categories:  categories,
	}
	return resp, nil
}
func getProduct(id int) (model.Product, []string, error) {
	var row model.Product
	var categoriesProduct []int
	var categories []string
	db := mysql.DB

	// 检查数据库连接是否有效
	if db == nil {
		return row, categories, fmt.Errorf("database connection is nil")
	}

	// 查询产品信息
	err := db.Model(&model.Product{}).Where("id=?", id).Find(&row).Error
	if err != nil {
		return row, categories, fmt.Errorf("failed to find product: %w", err)
	}

	// 查询产品类别ID
	err = db.Table("category_product").Select("category_id").Where("product_id= ?", id).Find(&categoriesProduct).Error
	if err != nil {
		return row, categories, fmt.Errorf("failed to find product categories: %w", err)
	}

	// 查询类别名称
	err = db.Table("categories").Select("name").Where("id in ?", categoriesProduct).Find(&categories).Error
	if err != nil {
		return row, categories, fmt.Errorf("failed to find categories: %w", err)
	}
	fmt.Printf("%+v\n%+v\n", row, categories)
	return row, categories, nil
}
