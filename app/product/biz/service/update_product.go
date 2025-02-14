package service

import (
	"context"
	"product/biz/dal/mysql"
	"product/biz/model"
	"product/kitex_gen/product"
)

type UpdateProductService struct {
	ctx context.Context
} // NewUpdateProductService new UpdateProductService
func NewUpdateProductService(ctx context.Context) *UpdateProductService {
	return &UpdateProductService{ctx: ctx}
}

// Run create note info
func (s *UpdateProductService) Run(req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	// Finish your business logic.
	resp = &product.UpdateProductResp{
		Success: true,
		Message: "ok",
	}
	temp := req.GetProduct()
	getProduct := &model.Product{
		Base:          model.Base{ID: int(temp.Id)},
		Name:          temp.Name,
		Price:         temp.Price,
		Description:   temp.Description,
		OriginalPrice: temp.OriginalPrice,
		Stock:         temp.Stock,
		Images:        temp.Images,
		Status:        int(temp.Status),
	}
	//收集标签
	var getCategory []string
	for _, cat := range temp.Categories {
		getCategory = append(getCategory, cat)
	}

	for _, category := range req.Product.Categories {
		//判断分类id是否存在
		category, err := model.GetByCategoryName(mysql.DB, s.ctx, category)
		if err != nil {
			return nil, err
		}
		//插入关联表
		newCategoryProduct := &model.CategoryProduct{
			ProductId:  int64(temp.GetId()),
			CategoryId: int64(category.ID),
		}
		err = model.CreateCPRelation(mysql.DB, newCategoryProduct)
		if err != nil {
			return nil, err
		}
	}
	//更新商品
	err = model.UpdateProduct(mysql.DB, getProduct)
	if err != nil {
		resp.Success = false
		resp.Message = err.Error()
		return resp, err
	}
	return resp, nil
}
