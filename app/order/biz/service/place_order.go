package service

import (
	"apis/rpc"
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/product"
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
	// TODO: 1. 验证库存
	_, err = rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
		Id: req.OrderItems.ProductId,
	})
	// TODO: 清空购物车

	// TODO: 2. 计算订单价格 (checkout RPC)

	// TODO: 3. 创建订单
	model.CreateOrder(mysql.DB, &model.Order{
		UserID:      int64(req.UserId),
		OrderStatus: 0,
	})

	return
}
