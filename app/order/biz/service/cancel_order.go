package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/constant"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
	"order/biz/dal/mysql"
	"order/biz/model"
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
	// 1. 确认订单状态
	od, err := model.GetOrder(mysql.DB, req.OrderId)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	if od.OrderStatus != constant.Canceled {
		klog.Error("订单状态错误", err)
		return nil, err
	}
	// 2. 取消订单
	if err = model.CancelOrder(mysql.DB, constant.Canceled, req.OrderId); err != nil {
		klog.Error(err)
		return nil, err
	}
	// 3. 返回结果
	resp = &order.CancelOrderResp{
		Success: true,
	}
	return resp, nil
}
