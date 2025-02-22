package service

import (
	"cart/biz/model"
	"context"
	"errors"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {

	// 参数检查
	if req.UserId == 0 {
		return nil, errors.New("user id is required")
	}

	// 查询 这个 user 的所有 商品
	var userCart cart.Cart // user 购物车中的 item
	userCart.UserId = req.UserId
	model.QueryItemsByUser(&userCart)

	return &cart.GetCartResp{Cart: &userCart}, nil
}
