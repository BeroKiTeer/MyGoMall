package service

import (
	"context"
	"fmt"
	"product/biz/dal/mysql"
	"product/biz/model"
	"product/kitex_gen/product"
)

type GetProductsBatchService struct {
	ctx context.Context
} // NewGetProductsBatchService new GetProductsBatchService
func NewGetProductsBatchService(ctx context.Context) *GetProductsBatchService {
	return &GetProductsBatchService{ctx: ctx}
}

// Run create note info
func (s *GetProductsBatchService) Run(req *product.GetProductsBatchReq) (resp *product.GetProductsBatchResp, err error) {
	// Finish your business logic.
	resp = &product.GetProductsBatchResp{}
	ids := make([]int, len(req.Ids))
	for i, v := range req.GetIds() {
		ids[i] = int(v) // 将 uint32 转换为 int
	}
	fmt.Println("ids", ids)
	products, categories, _ := model.GetProductsByIds(mysql.DB, ids)
	for i, v := range products {
		temp := &product.Product{
			Id:            uint32(v.ID),
			Name:          v.Name,
			Price:         v.Price,
			Stock:         v.Stock,
			Images:        v.Images,
			Description:   v.Description,
			OriginalPrice: v.OriginalPrice,
			Status:        uint32(v.Status),
			Categories:    categories[i],
		}
		resp.Products = append(resp.Products, temp)
	}
	return
}
