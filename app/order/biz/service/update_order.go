package service

import (
	"context"
	order "order/kitex_gen/order"
)

type UpdateOrderService struct {
	ctx context.Context
} // NewUpdateOrderService new UpdateOrderService
func NewUpdateOrderService(ctx context.Context) *UpdateOrderService {
	return &UpdateOrderService{ctx: ctx}
}

// Run create note info
func (s *UpdateOrderService) Run(req *order.UpdateOrderReq) (resp *order.UpdateOrderResp, err error) {
	// Finish your business logic.
	// TODO: 1. 验证用户有效性

	// TODO: 2. 修改订单信息

	// TODO: 3. 修改订单状态（订单有效期）
	return
}
