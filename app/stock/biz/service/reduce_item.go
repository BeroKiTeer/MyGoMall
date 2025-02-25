package service

import (
	"context"
	"errors"
	stock "github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	"stock/biz/dal/mysql"
	"stock/biz/model"
)

type ReduceItemService struct {
	ctx context.Context
} // NewReduceItemService new ReduceItemService
func NewReduceItemService(ctx context.Context) *ReduceItemService {
	return &ReduceItemService{ctx: ctx}
}

// Run create note info
func (s *ReduceItemService) Run(req *stock.ReduceItemReq) (resp *stock.ReduceItemResp, err error) {

	// 先看看有没有这么多商品
	quantity, err := model.CheckQuantity(mysql.DB, req.ProductId)

	// 数据库查询遇到了问题
	if err != nil {
		return &stock.ReduceItemResp{Success: false}, err
	}

	// 商品数量不足
	if quantity < req.Quantity {
		return &stock.ReduceItemResp{Success: false}, errors.New("商品数量不足！")
	}

	// 减少商品数量
	err = model.ReduceItem(mysql.DB, req.ProductId, req.Quantity)

	return &stock.ReduceItemResp{Success: true}, nil
}
