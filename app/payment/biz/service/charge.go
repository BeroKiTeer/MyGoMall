package service

import (
	"context"
	"errors"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/payment"
	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
	"payment/biz/dal/mysql"
	"payment/biz/model"
	"strconv"
	"time"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {

	// 模拟一个信用卡
	card := creditcard.Card{
		Number: req.CreditCard.CreditCardNumber,
		Cvv:    strconv.Itoa(int(req.CreditCard.CreditCardCvv)),
		Month:  strconv.Itoa(int(req.CreditCard.CreditCardExpirationMonth)),
		Year:   strconv.Itoa(int(req.CreditCard.CreditCardExpirationYear)),
	}

	// 验证信用卡是否有效
	err = card.Validate(true)
	if err != nil {
		return nil, errors.New("信用卡无效！" + err.Error())
	}

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
		Way:           "credit_card",
	})

	if err != nil {
		return nil, errors.New("数据库存储失败：" + err.Error())
	}

	return &payment.ChargeResp{TransactionId: transactionID.String()}, nil
}
