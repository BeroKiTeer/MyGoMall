package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	//TODO: 1. 验证用户有效性

	//TODO: 2. 查询用户的订单

	return
}
