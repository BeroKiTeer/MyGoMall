package service

import (
	"context"
	"encoding/json"
	order "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/klog"
	"order/biz/dal/mysql"
	"order/biz/model"
)

type ShowOrderDetailService struct {
	ctx context.Context
} // NewShowOrderDetailService new ShowOrderDetailService
func NewShowOrderDetailService(ctx context.Context) *ShowOrderDetailService {
	return &ShowOrderDetailService{ctx: ctx}
}

// Run create note info
func (s *ShowOrderDetailService) Run(req *order.ShowOrderDetailReq) (resp *order.ShowOrderDetailResp, err error) {
	// Finish your business logic.
	tx := mysql.DB.Begin()
	items, err := model.GetCachedItemsByOrderId(s.ctx, tx, req.OrderId)
	var orderitems []*order.OrderItem
	for _, item := range items {
		var ord_item model.OrderItem
		if err := json.Unmarshal([]byte(item), &ord_item); err != nil {
			klog.Error("订单反序列化失败:", err)
			return nil, err
		}
		if ord_item.OrderId != req.OrderId {
			return nil, err
		}
		orderitems = append(orderitems, &order.OrderItem{
			ProductId: ord_item.ProductId,
			Quantity:  int32(ord_item.Quantity),
			Cost:      ord_item.Price,
		})
	}
	if tx.Commit().Error != nil {
		// TODO: 其他服务事务补偿
		return nil, tx.Commit().Error
	}
	resp = &order.ShowOrderDetailResp{
		OrderItems: orderitems,
	}
	return resp, nil
}
