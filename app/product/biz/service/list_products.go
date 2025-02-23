package service

import (
	"context"
	"fmt"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/product"
	"product/biz/dal/mysql"
	"product/biz/model"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	resp = &product.ListProductsResp{}
	// Finish your business logic.
	fmt.Printf("请求页数:%#v\n", req.Page)
	fmt.Printf("商品数量:%#v\n", req.PageSize)
	fmt.Printf("搜索商品标签:%#v\n", req.CategoryName)
	products, categories, err := model.GetProductsByCategoryName(mysql.DB, int(req.Page), int(req.PageSize), req.CategoryName)
	for i, item := range products {
		pro := &product.Product{
			Id:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Images:      item.Images,
			Price:       item.Price,
			Categories:  categories[i],
		}
		resp.Products = append(resp.Products, pro)
	}
	return resp, nil
}
