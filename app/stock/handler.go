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

// ReserveItem implements the StockServiceImpl interface.
func (s *StockServiceImpl) ReserveItem(ctx context.Context, req *stock.ReserveItemReq) (resp *stock.ReserveItemResp, err error) {
	resp, err = service.NewReserveItemService(ctx).Run(req)

	return resp, err
}

// RecoverItem implements the StockServiceImpl interface.
func (s *StockServiceImpl) RecoverItem(ctx context.Context, req *stock.RecoverItemReq) (resp *stock.RecoverItemResp, err error) {
	resp, err = service.NewRecoverItemService(ctx).Run(req)

	return resp, err
}
