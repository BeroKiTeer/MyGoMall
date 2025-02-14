package service

import (
	cart "cart/kitex_gen/cart"
	"context"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// Finish your business logic.
	// TODO: 1. 参数检查

	// TODO: 2. 删除

	// TODO: 3. 返回。

	return
}
