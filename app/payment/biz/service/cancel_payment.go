package service

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/klog"
	"payment/biz/dal/mysql"
	"payment/biz/model"
)

type CancelPaymentService struct {
	ctx context.Context
}

// NewCancelPaymentService new CancelPaymentService
func NewCancelPaymentService(ctx context.Context) *CancelPaymentService {
	return &CancelPaymentService{ctx: ctx}
}

// Run create note info
func (s *CancelPaymentService) Run(req *payment.CancelReq) (resp *payment.CancelResp, err error) {
	cancelResp := &payment.CancelResp{
		Status: "error",
	}
	//查询订单信息
	pay, err := model.QueryPayment(mysql.DB, int(req.Id))
	if err != nil {
		klog.Error("QueryPayment接口未响应，查询订单信息失败", err)
		return cancelResp, err
	}
	// 判断订单状态
	if pay.Status == 0 {
		//待支付，将订单状态更新为取消 4
		pay.Status = 4
		err := model.UpdatePaymentStatus(mysql.DB, pay)
		if err != nil {
			klog.Error("UpdatePaymentStatus接口未响应，订单更新失败", err)
			cancelResp.Status = "订单更新失败，请重试"
			return cancelResp, err
		}
	} else if pay.Status == 1 {
		//---已支付，先退款，再取消
		//暂时使用待支付取消逻辑
		pay.Status = 4
		err := model.UpdatePaymentStatus(mysql.DB, pay)
		if err != nil {
			klog.Error("UpdatePaymentStatus接口未响应，订单更新失败", err)
			cancelResp.Status = "订单更新失败，请重试"
			return cancelResp, err
		}
	} else {
		//其他情况
		cancelResp.Status = "订单已取消或已退款，请勿重复操作"
		return cancelResp, nil
	}
	return cancelResp, nil
}
