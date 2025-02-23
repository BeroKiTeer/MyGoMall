package main

import (
	"context"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	"stock/biz/service"
)

// StockServiceImpl implements the last service interface defined in the IDL.
type StockServiceImpl struct{}

// ReduceItem implements the StockServiceImpl interface.
func (s *StockServiceImpl) ReduceItem(ctx context.Context, req *stock.ReduceItemReq) (resp *stock.ReduceItemResp, err error) {
	resp, err = service.NewReduceItemService(ctx).Run(req)

	return resp, err
}

// CheckItem implements the StockServiceImpl interface.
func (s *StockServiceImpl) CheckItem(ctx context.Context, req *stock.CheckItemReq) (resp *stock.CheckItemResp, err error) {
	resp, err = service.NewCheckItemService(ctx).Run(req)

	return resp, err
}
