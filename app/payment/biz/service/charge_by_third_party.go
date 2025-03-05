package service

import (
	"context"
	"errors"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/payment"
	"github.com/cloudwego/kitex/pkg/klog"
	"payment/biz/dal/mysql"
	mq "payment/biz/dal/rabbitmq"
	"payment/biz/model"
	"time"
)

type ChargeByThirdPartyService struct {
	ctx context.Context
} // NewChargeByThirdPartyService new ChargeByThirdPartyService
func NewChargeByThirdPartyService(ctx context.Context) *ChargeByThirdPartyService {
	return &ChargeByThirdPartyService{ctx: ctx}
}

// Run create note info
func (s *ChargeByThirdPartyService) Run(req *payment.ChargeByThirdPartyReq) (resp *payment.ChargeByThirdPartyResp, err error) {
	//发送消息到队列
	klog.Info("send.....")
	err = mq.CardPaymentProducer.Send(&mq.CardPayment{
		OrderID:       req.OrderId,
		TransactionID: req.TransactionId,
		Success:       true,
	})
	klog.Info("send.....")
	if err != nil {
		klog.Errorf("消息发送失败%v", err)
		return nil, err
	}
	err = model.CreatePayment(mysql.DB, &model.Payment{
		UserID:        req.UserId,
		OrderID:       req.OrderId,
		TransactionID: req.TransactionId,
		Amount:        req.Amount,
		PayAt:         time.Now(),
		Way:           req.Way,
	})

	if err != nil {
		klog.Error("数据库存储失败", err)
		return nil, errors.New("数据库存储失败：" + err.Error())
	}

	return &payment.ChargeByThirdPartyResp{TransactionId: req.TransactionId}, nil

}
