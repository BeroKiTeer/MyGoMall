package auth

import (
	"context"
	auth "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/auth"

	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/auth/authservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() authservice.Client
	Service() string
	DeliverTokenByRPC(ctx context.Context, Req *auth.DeliverTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error)
	VerifyTokenByRPC(ctx context.Context, Req *auth.VerifyTokenReq, callOptions ...callopt.Option) (r *auth.VerifyResp, err error)
	RefreshToken(ctx context.Context, Req *auth.RefreshTokenReq, callOptions ...callopt.Option) (r *auth.RefreshTokenResp, err error)
	DecodeToken(ctx context.Context, Req *auth.DecodeTokenReq, callOptions ...callopt.Option) (r *auth.DecodeTokenResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := authservice.NewClient(dstService, opts...)
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
	kitexClient authservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() authservice.Client {
	return c.kitexClient
}

func (c *clientImpl) DeliverTokenByRPC(ctx context.Context, Req *auth.DeliverTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error) {
	return c.kitexClient.DeliverTokenByRPC(ctx, Req, callOptions...)
}

func (c *clientImpl) VerifyTokenByRPC(ctx context.Context, Req *auth.VerifyTokenReq, callOptions ...callopt.Option) (r *auth.VerifyResp, err error) {
	return c.kitexClient.VerifyTokenByRPC(ctx, Req, callOptions...)
}

func (c *clientImpl) RefreshToken(ctx context.Context, Req *auth.RefreshTokenReq, callOptions ...callopt.Option) (r *auth.RefreshTokenResp, err error) {
	return c.kitexClient.RefreshToken(ctx, Req, callOptions...)
}

func (c *clientImpl) DecodeToken(ctx context.Context, Req *auth.DecodeTokenReq, callOptions ...callopt.Option) (r *auth.DecodeTokenResp, err error) {
	return c.kitexClient.DecodeToken(ctx, Req, callOptions...)
}
