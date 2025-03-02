package service

import (
	"cart/biz/model"
	"context"
	"errors"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/klog"
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
		klog.Error("未输入用户id", err)
		return nil, errors.New("empty user id")
	}

	// 检查商品是否已存在在购物车
	var targetItemQuantity int32 = -1
	err = model.CheckItemsByUser(req.UserId, &targetItemQuantity)

	if err != nil {
		klog.Error("未查询到商品", err)
		return nil, err
	}

	// 删除
	if targetItemQuantity != -1 {
		err = model.EmptyCart(req.UserId)
	}
	return &cart.EmptyCartResp{}, err
}
