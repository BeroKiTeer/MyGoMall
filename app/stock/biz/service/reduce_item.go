package service

import (
	"context"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
)

type ReduceItemService struct {
	ctx context.Context
} // NewReduceItemService new ReduceItemService
func NewReduceItemService(ctx context.Context) *ReduceItemService {
	return &ReduceItemService{ctx: ctx}
}

// Run create note info
func (s *ReduceItemService) Run(req *stock.ReduceItemReq) (resp *stock.ReduceItemResp, err error) {
	// Finish your business logic.

	return
}
