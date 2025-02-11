package service

import (
	"context"
	product "product/kitex_gen/product"
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

	return
}
