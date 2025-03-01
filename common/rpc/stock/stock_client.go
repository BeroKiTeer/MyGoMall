package stock

import (
	"context"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"

	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock/stockservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() stockservice.Client
	Service() string
	ReduceItem(ctx context.Context, Req *stock.ReduceItemReq, callOptions ...callopt.Option) (r *stock.ReduceItemResp, err error)
	CheckItem(ctx context.Context, Req *stock.CheckItemReq, callOptions ...callopt.Option) (r *stock.CheckItemResp, err error)
	ReserveItem(ctx context.Context, Req *stock.ReserveItemReq, callOptions ...callopt.Option) (r *stock.ReserveItemResp, err error)
	RecoverItem(ctx context.Context, Req *stock.RecoverItemReq, callOptions ...callopt.Option) (r *stock.RecoverItemResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := stockservice.NewClient(dstService, opts...)
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
	kitexClient stockservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() stockservice.Client {
	return c.kitexClient
}

func (c *clientImpl) ReduceItem(ctx context.Context, Req *stock.ReduceItemReq, callOptions ...callopt.Option) (r *stock.ReduceItemResp, err error) {
	return c.kitexClient.ReduceItem(ctx, Req, callOptions...)
}

func (c *clientImpl) CheckItem(ctx context.Context, Req *stock.CheckItemReq, callOptions ...callopt.Option) (r *stock.CheckItemResp, err error) {
	return c.kitexClient.CheckItem(ctx, Req, callOptions...)
}

func (c *clientImpl) ReserveItem(ctx context.Context, Req *stock.ReserveItemReq, callOptions ...callopt.Option) (r *stock.ReserveItemResp, err error) {
	return c.kitexClient.ReserveItem(ctx, Req, callOptions...)
}

func (c *clientImpl) RecoverItem(ctx context.Context, Req *stock.RecoverItemReq, callOptions ...callopt.Option) (r *stock.RecoverItemResp, err error) {
	return c.kitexClient.RecoverItem(ctx, Req, callOptions...)
}
