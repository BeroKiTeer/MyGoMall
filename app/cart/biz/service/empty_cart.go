package service

import (
	"cart/biz/dal/mysql"
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
	mysql.DB.Table("carts").
		Select("SUM(quantity)").
		Where("user_id = ?", req.UserId).
		Group("user_id").Scan(&targetItemQuantity)

	// 删除
	if targetItemQuantity != -1 {
		mysql.DB.Table("carts").
			Where("user_id = ?", req.UserId).
			Delete(&req.UserId)
	}

	return
}
