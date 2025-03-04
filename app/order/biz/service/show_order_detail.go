package service

import (
	"context"
	order "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
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
	items, err := model.SelectOrderItemsById(tx, req.OrderId)
	var orderitems []*order.OrderItem
	for _, item := range items {
		if item.OrderId != req.OrderId {
			return nil, err
		}
		orderitems = append(orderitems, &order.OrderItem{
			ProductId: item.ProductId,
			Quantity:  int32(item.Quantity),
			Cost:      item.Price,
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
