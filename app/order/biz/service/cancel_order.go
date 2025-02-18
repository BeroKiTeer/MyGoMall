package service

import (
	"context"
	order "order/kitex_gen/order"
)

type CancelOrderService struct {
	ctx context.Context
} // NewCancelOrderService new CancelOrderService
func NewCancelOrderService(ctx context.Context) *CancelOrderService {
	return &CancelOrderService{ctx: ctx}
}

// Run create note info
func (s *CancelOrderService) Run(req *order.CancelOrderReq) (resp *order.CancelOrderResp, err error) {
	// Finish your business logic.

	return
}
