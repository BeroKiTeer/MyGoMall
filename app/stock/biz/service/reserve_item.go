package service

import (
	"context"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
)

type ReserveItemService struct {
	ctx context.Context
} // NewReserveItemService new ReserveItemService
func NewReserveItemService(ctx context.Context) *ReserveItemService {
	return &ReserveItemService{ctx: ctx}
}

// Run create note info
func (s *ReserveItemService) Run(req *stock.ReserveItemReq) (resp *stock.ReserveItemResp, err error) {
	// Finish your business logic.

	return
}
