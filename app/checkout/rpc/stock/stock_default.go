package stock

import (
	"context"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ReduceItem(ctx context.Context, req *stock.ReduceItemReq, callOptions ...callopt.Option) (resp *stock.ReduceItemResp, err error) {
	resp, err = defaultClient.ReduceItem(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ReduceItem call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CheckItem(ctx context.Context, req *stock.CheckItemReq, callOptions ...callopt.Option) (resp *stock.CheckItemResp, err error) {
	resp, err = defaultClient.CheckItem(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CheckItem call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ReserveItem(ctx context.Context, req *stock.ReserveItemReq, callOptions ...callopt.Option) (resp *stock.ReserveItemResp, err error) {
	resp, err = defaultClient.ReserveItem(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ReserveItem call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
