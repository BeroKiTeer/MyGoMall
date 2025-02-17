package service

import (
	"cart/biz/model"
	cart "cart/kitex_gen/cart"
	"context"
	"errors"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {

	// 参数检查
	if req.UserId == 0 {
		return nil, errors.New("empty user id")
	}

	// 检查商品是否已存在在购物车
	var targetItemQuantity int32 = -1
	model.CheckItemsByUser(req.UserId, &targetItemQuantity)

	// 删除
	if targetItemQuantity != -1 {
		model.EmptyCart(req.UserId)
	}

	return
}
