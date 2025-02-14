package service

import (
	cart "cart/kitex_gen/cart"
	"context"
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

var db, err = gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {

	// 参数检查
	if req.UserId == 0 {
		return nil, errors.New("user id is required")
	}

	// 查询 这个 user 的所有 商品
	var userCart cart.Cart // user 购物车中的 item
	db.Select("product_id", "quantity").Find(&userCart.Items)

	return &cart.GetCartResp{Cart: &userCart}, nil
}
