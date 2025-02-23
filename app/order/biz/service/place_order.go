package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
)

type PlaceOrderService struct {
	ctx context.Context
} // NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	// TODO: 1. 验证token的有效性

	// TODO: 2. 验证库存

	// TODO: 3. 计算订单价格 (checkout RPC)

	// TODO: 4. 创建订单

	return
}
