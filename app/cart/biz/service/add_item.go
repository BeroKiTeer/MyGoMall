package service

import (
	cart "cart/kitex_gen/cart"
	"context"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	// TODO: 1.参数验证

	// TODO: 2.检查商品是否存在（RPC）

	// TODO: 3.检查商品库存是否足够（可选）

	// TODO: 4.检查商品是否已存在在购物车

	// TODO: 5.将商品添加到购物车

	// TODO: 6.持久化存储

	return
}
