package main

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	"order/biz/service"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	resp, err = service.NewPlaceOrderService(ctx).Run(req)

	return resp, err
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	resp, err = service.NewListOrderService(ctx).Run(req)

	return resp, err
}

// MarkOrderPaid implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	resp, err = service.NewMarkOrderPaidService(ctx).Run(req)

	return resp, err
}

// UpdateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrder(ctx context.Context, req *order.UpdateOrderReq) (resp *order.UpdateOrderResp, err error) {
	resp, err = service.NewUpdateOrderService(ctx).Run(req)

	return resp, err
}

// CancelOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CancelOrder(ctx context.Context, req *order.CancelOrderReq) (resp *order.CancelOrderResp, err error) {
	resp, err = service.NewCancelOrderService(ctx).Run(req)

	return resp, err
}

// ShowOrderDetail implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ShowOrderDetail(ctx context.Context, req *order.ShowOrderDetailReq) (resp *order.ShowOrderDetailResp, err error) {
	resp, err = service.NewShowOrderDetailService(ctx).Run(req)

	return resp, err
}
