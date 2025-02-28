package order

import (
	"context"
	order "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"

	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() orderservice.Client
	Service() string
	PlaceOrder(ctx context.Context, Req *order.PlaceOrderReq, callOptions ...callopt.Option) (r *order.PlaceOrderResp, err error)
	ListOrder(ctx context.Context, Req *order.ListOrderReq, callOptions ...callopt.Option) (r *order.ListOrderResp, err error)
	MarkOrderPaid(ctx context.Context, Req *order.MarkOrderPaidReq, callOptions ...callopt.Option) (r *order.MarkOrderPaidResp, err error)
	UpdateOrder(ctx context.Context, Req *order.UpdateOrderReq, callOptions ...callopt.Option) (r *order.UpdateOrderResp, err error)
	CancelOrder(ctx context.Context, Req *order.CancelOrderReq, callOptions ...callopt.Option) (r *order.CancelOrderResp, err error)
	ShowOrderDetail(ctx context.Context, Req *order.ShowOrderDetailReq, callOptions ...callopt.Option) (r *order.ShowOrderDetailResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := orderservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient orderservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() orderservice.Client {
	return c.kitexClient
}

func (c *clientImpl) PlaceOrder(ctx context.Context, Req *order.PlaceOrderReq, callOptions ...callopt.Option) (r *order.PlaceOrderResp, err error) {
	return c.kitexClient.PlaceOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) ListOrder(ctx context.Context, Req *order.ListOrderReq, callOptions ...callopt.Option) (r *order.ListOrderResp, err error) {
	return c.kitexClient.ListOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) MarkOrderPaid(ctx context.Context, Req *order.MarkOrderPaidReq, callOptions ...callopt.Option) (r *order.MarkOrderPaidResp, err error) {
	return c.kitexClient.MarkOrderPaid(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateOrder(ctx context.Context, Req *order.UpdateOrderReq, callOptions ...callopt.Option) (r *order.UpdateOrderResp, err error) {
	return c.kitexClient.UpdateOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) CancelOrder(ctx context.Context, Req *order.CancelOrderReq, callOptions ...callopt.Option) (r *order.CancelOrderResp, err error) {
	return c.kitexClient.CancelOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) ShowOrderDetail(ctx context.Context, Req *order.ShowOrderDetailReq, callOptions ...callopt.Option) (r *order.ShowOrderDetailResp, err error) {
	return c.kitexClient.ShowOrderDetail(ctx, Req, callOptions...)
}
