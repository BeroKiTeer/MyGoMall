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
	// TODO: 1. 验证用户有效性

	// TODO: 2. 库存解锁

	// TODO: 3. 订单状态改为已取消

	return
}
