package service

import (
	"context"
	"encoding/json"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	"order/biz/dal/mysql"
	"order/biz/model"
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
	resp = &order.ListOrderResp{}

	orders, err := model.GetCachedListedOrdersById(s.ctx, mysql.DB, int64(req.UserId))
	for _, item := range orders {
		var ord order.Order
		json.Unmarshal([]byte(item), &ord)

		ord_itms, _ := model.GetCachedItemsByOrderId(s.ctx, mysql.DB, ord.OrderId)
		for _, pro := range ord_itms {
			var item order.OrderItem
			json.Unmarshal([]byte(pro), &item)

			ord.OrderItems = append(ord.OrderItems, &item)
		}

		resp.Orders = append(resp.Orders, &ord)
	}
	return resp, nil

}
