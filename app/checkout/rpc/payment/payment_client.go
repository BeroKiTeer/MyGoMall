package payment

import (
	"context"
	payment "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/payment"

	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() paymentservice.Client
	Service() string
	Charge(ctx context.Context, Req *payment.ChargeReq, callOptions ...callopt.Option) (r *payment.ChargeResp, err error)
	CancelPayment(ctx context.Context, Req *payment.CancelReq, callOptions ...callopt.Option) (r *payment.CancelResp, err error)
	ChargeByThirdParty(ctx context.Context, Req *payment.ChargeByThirdPartyReq, callOptions ...callopt.Option) (r *payment.ChargeByThirdPartyResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := paymentservice.NewClient(dstService, opts...)
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
	kitexClient paymentservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() paymentservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Charge(ctx context.Context, Req *payment.ChargeReq, callOptions ...callopt.Option) (r *payment.ChargeResp, err error) {
	return c.kitexClient.Charge(ctx, Req, callOptions...)
}

func (c *clientImpl) CancelPayment(ctx context.Context, Req *payment.CancelReq, callOptions ...callopt.Option) (r *payment.CancelResp, err error) {
	return c.kitexClient.CancelPayment(ctx, Req, callOptions...)
}

func (c *clientImpl) ChargeByThirdParty(ctx context.Context, Req *payment.ChargeByThirdPartyReq, callOptions ...callopt.Option) (r *payment.ChargeByThirdPartyResp, err error) {
	return c.kitexClient.ChargeByThirdParty(ctx, Req, callOptions...)
}
