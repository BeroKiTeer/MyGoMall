package service

import (
	"context"
	"product/biz/dal/mysql"
	"product/biz/model"
	product "product/kitex_gen/product"
)

type AddProductService struct {
	ctx context.Context
} // NewAddProductService new AddProductService
func NewAddProductService(ctx context.Context) *AddProductService {
	return &AddProductService{ctx: ctx}
}

// Run create note info
// 添加商品
func (s *AddProductService) Run(req *product.AddProductReq) (resp *product.AddProductResp, err error) {
	// Finish your business logic.

	newProduct := &model.Product{
		CategoryId:    int(req.CategoryId),
		Name:          req.Name,
		Description:   req.Description,
		Price:         float64(req.Price),
		OriginalPrice: float64(req.OriginalPrice),
		Stock:         int(req.Stock),
		Images:        req.Images,
		Status:        int(req.Status),
	}
	err = model.AddProduct(mysql.DB, newProduct)
	if err != nil {
		return nil, err
	}
	return &product.AddProductResp{ProductId: int32(newProduct.ID)}, err
}
