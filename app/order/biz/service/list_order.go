package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	"order/biz/model"
	"product/biz/dal/mysql"
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

	orders, err := model.GetOrdersByUserID(mysql.DB, int64(req.UserId))
	for _, item := range orders {
		ord := &order.Order{
			OrderId: item.ID,
			UserId:  uint32(item.UserID),
			Address: &order.Address{
				TelephoneNumber: item.PhoneNumber,
				StreetAddress:   item.ShippingAddress,
				Name:            item.RecipientName,
			},
			OrderItems: []*order.OrderItem{},
		}
		ord_itms, _ := model.GetOrderItemByOrderID(mysql.DB, item.ID)
		for _, pro := range ord_itms {
			orderItem := &order.OrderItem{
				ProductId: pro.ProductId,
				Quantity:  int32(pro.Quantity),
				Cost:      pro.Price,
			}
			ord.OrderItems = append(ord.OrderItems, orderItem)
		}

		resp.Orders = append(resp.Orders, ord)
	}
	return resp, nil

}
