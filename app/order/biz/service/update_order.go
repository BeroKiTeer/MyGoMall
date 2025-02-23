package service

import (
	"context"
	"fmt"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	"order/biz/model"
	"product/biz/dal/mysql"
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

	resp = &order.UpdateOrderResp{
		Success: true,
	}
	// 从数据库中获取订单
	ord, err := model.GetOrder(mysql.DB, req.OrderId)
	if err != nil {
		resp.Success = false
		return resp, fmt.Errorf("failed to get order: %w", err)
	}

	// 更新订单信息
	if req.Address != nil {
		ord.ShippingAddress = req.Address.StreetAddress
		ord.RecipientName = req.Address.Name
		ord.PhoneNumber = req.Address.TelephoneNumber
	}

	// 将更新后的订单保存回数据库
	err = model.UpdateOrder(mysql.DB, &ord)
	if err != nil {
		resp.Success = false
		return resp, fmt.Errorf("failed to update order: %w", err)
	}

	return resp, nil
}
