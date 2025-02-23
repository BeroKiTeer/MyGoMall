package service

import (
	"context"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
)

type CheckItemService struct {
	ctx context.Context
} // NewCheckItemService new CheckItemService
func NewCheckItemService(ctx context.Context) *CheckItemService {
	return &CheckItemService{ctx: ctx}
}

// Run create note info
func (s *CheckItemService) Run(req *stock.CheckItemReq) (resp *stock.CheckItemResp, err error) {
	// Finish your business logic.

	return
}
