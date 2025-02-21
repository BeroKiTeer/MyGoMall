package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	"order/biz/model"
	"product/biz/dal/mysql"
	"strconv"
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
			OrderId: strconv.Itoa(item.ID),
			UserId:  uint32(item.UserID),
			//这里还有问题，比如OrderStruct里面的OrderItems是*OrderItems类型，我看看怎么弄
		}
		resp.Orders = append(resp.Orders, ord)
	}
	return resp, nil

}
