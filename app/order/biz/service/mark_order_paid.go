package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
)

type MarkOrderPaidService struct {
	ctx context.Context
} // NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// Finish your business logic.

	return
}
