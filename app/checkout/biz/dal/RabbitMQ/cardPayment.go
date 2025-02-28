package rabbitMq

import (
	"checkout/conf"
	"github.com/cloudwego/hertz/pkg/common/json"
)

var (
	CardPaymentProducer, _ = NewPaymentProducer(MQConfig{
		Exchange:     conf.GetConf().RabbitMQ.Payments.Methods["creditCard"].Exchange,
		QueueName:    conf.GetConf().RabbitMQ.Payments.Methods["creditCard"].Queue,
		RoutingKey:   conf.GetConf().RabbitMQ.Payments.Methods["creditCard"].RoutingKey,
		ExchangeType: conf.GetConf().RabbitMQ.Payments.Methods["creditCard"].ExchangeType,
		MqURL:        conf.GetConf().RabbitMQ.MqURL,
	})
)

// 信用卡支付
type CardPayment struct {
	OrderID     string `json:"order_id"`
	Amount      int64  `json:"amount"`
	CallbackURL string `json:"callback_url"`
}

func (c *CardPayment) GetRoutingKey() string {
	return CardPaymentProducer.config.RoutingKey // 专属路由键
}

func (c *CardPayment) Marshal() ([]byte, error) {
	return json.Marshal(c)
}
