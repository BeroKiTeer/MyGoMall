package service

import (
	"context"
	"errors"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/payment"
	"github.com/google/uuid"
	"payment/biz/dal/mysql"
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

	// 生成一个交易 ID （UUID）
	transactionID, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.New("UUID生成失败：" + err.Error())
	}

	err = model.CreatePayment(mysql.DB, &model.Payment{
		UserID:        req.UserId,
		OrderID:       req.OrderId,
		TransactionID: transactionID.String(),
		Amount:        req.Amount,
		PayAt:         time.Now(),
		Way:           req.Way,
	})

	if err != nil {
		return nil, errors.New("数据库存储失败：" + err.Error())
	}

	return &payment.ChargeByThirdPartyResp{TransactionId: transactionID.String()}, nil

}
