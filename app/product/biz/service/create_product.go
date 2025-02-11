package service

import (
	"context"
	"product/biz/dal/mysql"
	"product/biz/model"
	product "product/kitex_gen/product"
)

type CreateProductService struct {
	ctx context.Context
} // NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// Run create note info
func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	// Finish your business logic.

	newProduct := &model.Product{
		Name:          req.Product.Name,
		Description:   req.Product.Description,
		Price:         float32(req.Product.Price),
		OriginalPrice: float32(req.Product.OriginalPrice),
		Stock:         req.Product.Stock,
		Images:        req.Product.Images,
		Status:        int(req.Product.Status),
	}
	//TODO:需要添加事务处理
	//插入product表and 关联表
	err = model.CreateProduct(mysql.DB, newProduct)
	if err != nil {
		return nil, err
	}
	//处理分类id
	for _, category := range req.Product.Categories {
		//判断分类id是否存在
		category, err := model.GetByCategoryName(mysql.DB, s.ctx, category)
		if err != nil {
			return nil, err
		}
		//插入关联表
		newCategoryProduct := &model.CategoryProduct{
			ProductId:  int64(newProduct.ID),
			CategoryId: int64(category.ID),
		}
		err = model.CreateCPRelation(mysql.DB, newCategoryProduct)
		if err != nil {
			return nil, err
		}
	}
	return &product.CreateProductResp{ProductId: uint32(newProduct.ID)}, err
}
