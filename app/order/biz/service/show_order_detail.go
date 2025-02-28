package service

import (
	"context"
	order "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
)

type ShowOrderDetailService struct {
	ctx context.Context
} // NewShowOrderDetailService new ShowOrderDetailService
func NewShowOrderDetailService(ctx context.Context) *ShowOrderDetailService {
	return &ShowOrderDetailService{ctx: ctx}
}

// Run create note info
func (s *ShowOrderDetailService) Run(req *order.ShowOrderDetailReq) (resp *order.ShowOrderDetailResp, err error) {
	// Finish your business logic.

	return
}
