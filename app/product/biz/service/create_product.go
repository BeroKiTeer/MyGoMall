package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/product"
	"product/biz/dal/mysql"
	"product/biz/model"
)

type CreateProductService struct {
	ctx context.Context
} // NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// Run create note info
func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	// 开始事务
	tx := mysql.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	newProduct := &model.Product{
		Name:          req.Product.Name,
		Description:   req.Product.Description,
		Price:         req.Product.Price,
		OriginalPrice: req.Product.OriginalPrice,
		Stock:         req.Product.Stock,
		Images:        req.Product.Images,
		Status:        int(req.Product.Status),
	}

	// 插入 product 表
	result := tx.Create(newProduct)
	if result.Error != nil {
		// 发生错误时回滚事务
		tx.Rollback()
		return nil, result.Error
	}

	// 处理分类 id
	for _, categoryName := range req.Product.Categories {
		// 判断分类 id 是否存在
		var category model.Categories
		result = tx.Where("name = ?", categoryName).First(&category)
		if result.Error != nil {
			// 发生错误时回滚事务
			tx.Rollback()
			return nil, result.Error
		}
		// 插入关联表
		newCategoryProduct := &model.CategoryProduct{
			ProductId:  newProduct.ID,
			CategoryId: category.ID,
		}
		result = tx.Create(newCategoryProduct)
		if result.Error != nil {
			// 发生错误时回滚事务
			tx.Rollback()
			return nil, result.Error
		}
	}
	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, err
	}
	return &product.CreateProductResp{ProductId: newProduct.ID}, nil
}
