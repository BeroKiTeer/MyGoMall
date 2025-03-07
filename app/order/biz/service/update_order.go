package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
	"order/biz/dal/mysql"
	"order/biz/model"
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
	var Order model.Order
	ordCached, err := model.GetCachedOrderById(s.ctx, mysql.DB, req.OrderId)
	if err != nil {
		resp.Success = false
		return resp, fmt.Errorf("failed to get order: %w", err)
	}
	if err := json.Unmarshal([]byte(ordCached), &Order); err != nil {
		klog.Error("用户反序列化失败:", err)
		return nil, err
	}
	// 更新订单信息
	if req.Address != nil {
		Order.RecipientName = req.Address.Name
		Order.PhoneNumber = req.Address.TelephoneNumber
	}

	// 将更新后的订单保存回数据库
	err = model.UpdateOrder(mysql.DB, &Order)
	if err != nil {
		resp.Success = false
		return resp, fmt.Errorf("failed to update order: %w", err)
	}

	return resp, nil
}
