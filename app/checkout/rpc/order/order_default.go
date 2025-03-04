package order

import (
	"context"
	order "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func PlaceOrder(ctx context.Context, req *order.PlaceOrderReq, callOptions ...callopt.Option) (resp *order.PlaceOrderResp, err error) {
	resp, err = defaultClient.PlaceOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "PlaceOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListOrder(ctx context.Context, req *order.ListOrderReq, callOptions ...callopt.Option) (resp *order.ListOrderResp, err error) {
	resp, err = defaultClient.ListOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func MarkOrderPaid(ctx context.Context, req *order.MarkOrderPaidReq, callOptions ...callopt.Option) (resp *order.MarkOrderPaidResp, err error) {
	resp, err = defaultClient.MarkOrderPaid(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "MarkOrderPaid call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateOrder(ctx context.Context, req *order.UpdateOrderReq, callOptions ...callopt.Option) (resp *order.UpdateOrderResp, err error) {
	resp, err = defaultClient.UpdateOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CancelOrder(ctx context.Context, req *order.CancelOrderReq, callOptions ...callopt.Option) (resp *order.CancelOrderResp, err error) {
	resp, err = defaultClient.CancelOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CancelOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ShowOrderDetail(ctx context.Context, req *order.ShowOrderDetailReq, callOptions ...callopt.Option) (resp *order.ShowOrderDetailResp, err error) {
	resp, err = defaultClient.ShowOrderDetail(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ShowOrderDetail call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
