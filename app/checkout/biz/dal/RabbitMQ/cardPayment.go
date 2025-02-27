package rabbitMq

import (
	"github.com/cloudwego/hertz/pkg/common/json"
)

var (
	CardPaymentProducer *PaymentProducer
)

// 信用卡支付
type CardPayment struct {
	OrderID     string `json:"order_id"`
	Amount      int64  `json:"amount"`
	CallbackURL string `json:"callback_url"`
}

func (c *CardPayment) GetRoutineKey() string {
	return CardPaymentProducer.config.RoutineKey // 专属路由键
}

func (c *CardPayment) Marshal() ([]byte, error) {
	return json.Marshal(c)
}
