package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	"order/biz/dal/mysql"
	"order/biz/model"
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
	// TODO: 1. 验证库存 （库存服务RPC调用）

	// TODO: 2. 清空购物车 （购物车服务RPC调用）

	// TODO: 3. 计算订单价格 (checkout RPC)

	// TODO: 4. 创建订单
	model.CreateOrder(mysql.DB, &model.Order{
		UserID:      int64(req.UserId),
		OrderStatus: 0,
	})

	return
}
