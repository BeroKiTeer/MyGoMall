package service

import (
	"context"
	"encoding/json"
	"github.com/BeroKiTeer/MyGoMall/common/constant"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
	"order/biz/dal/mysql"
	"order/biz/dal/redis"
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
	var Order model.Order
	od, err := model.GetCachedOrderById(s.ctx, mysql.DB, req.OrderId)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	if err := json.Unmarshal([]byte(od), &Order); err != nil {
		klog.Error("用户反序列化失败:", err)
		return nil, err
	}
	if Order.OrderStatus != constant.Canceled {
		klog.Error("订单状态错误", err)
		return nil, err
	}
	// 2. 取消订单
	redis.RedisClusterClient.Del(s.ctx, req.OrderId)
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
