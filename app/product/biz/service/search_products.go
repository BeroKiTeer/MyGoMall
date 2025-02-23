package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/product"
	"product/biz/dal/mysql"
	"product/biz/model"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	products, categories, err := model.GetProductByName(mysql.DB, req.GetName())
	if err != nil {
		return nil, err
	}
	resp = &product.SearchProductsResp{}
	for i, item := range products {
		pro := &product.Product{
			Id:          uint32(item.ID),
			Name:        item.Name,
			Description: item.Description,
			Images:      item.Images,
			Price:       item.Price,
			Categories:  categories[i],
		}
		resp.Results = append(resp.Results, pro)
	}
	return resp, err
}
