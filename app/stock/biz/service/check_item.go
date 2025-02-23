package service

import (
	"context"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	"stock/biz/model"
)

type CheckItemService struct {
	ctx context.Context
} // NewCheckItemService new CheckItemService
func NewCheckItemService(ctx context.Context) *CheckItemService {
	return &CheckItemService{ctx: ctx}
}

// Run create note info
func (s *CheckItemService) Run(req *stock.CheckItemReq) (resp *stock.CheckItemResp, err error) {

	// 查询商品剩余
	quantity, err := model.CheckQuantity(req.ProductId)

	return &stock.CheckItemResp{Quantity: quantity}, err
}
