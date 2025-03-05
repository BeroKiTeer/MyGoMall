package rabbitmq

import (
	"encoding/json"
	"payment/conf"
)

var (
	CardPaymentProducer, _ = NewPaymentProducer(MQConfig{
		Exchange:     conf.GetConf().RabbitMQ.Payments.Methods["creditCard"].Exchange,
		QueueName:    conf.GetConf().RabbitMQ.Payments.Methods["creditCard"].Queue,
		RoutingKey:   conf.GetConf().RabbitMQ.Payments.Methods["creditCard"].RoutingKey,
		ExchangeType: conf.GetConf().RabbitMQ.Payments.Methods["creditCard"].ExchangeType,
		MqURL:        conf.GetConf().RabbitMQ.RabbitMQURL,
	})
)

// 信用卡支付
type CardPayment struct {
	OrderID       string `json:"order_id"`
	TransactionID string `json:"transaction_id"`
	Success       bool   `json:"success"`
}

func (c *CardPayment) GetRoutingKey() string {
	return CardPaymentProducer.config.RoutingKey // 专属路由键
}
func (c *CardPayment) Marshal() ([]byte, error) {
	return json.Marshal(c)
}
