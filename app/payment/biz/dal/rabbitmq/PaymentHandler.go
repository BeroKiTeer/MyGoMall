package rabbitmq

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"net/http"
	"payment/conf"
)

type PaymentReq struct {
	OrderId string `json:"order_id"`
	Amount  string `json:"amount"`
	URL     bool   `json:"url"` //支付状态（成功还是失败）
}

type PaymentHandler struct {
	req PaymentReq
}

func (p *PaymentHandler) Marshal() ([]byte, error) {
	return json.Marshal(p.req)
}

func (p *PaymentHandler) Unmarshal(data []byte, i interface{}) error {

	err := json.Unmarshal(data, &p.req)
	if err != nil {
		return err
	}
	return nil

}

func (p *PaymentHandler) GetQueueName() (string, error) {
	config, err := conf.GetQueueConfig("payment_processor")
	if err != nil {
		return "", err
	}
	return config.Queue, nil
}

func (p *PaymentHandler) ProcessMessage(ctx context.Context, msg amqp.Delivery) error {
	var req PaymentReq
	err := p.Unmarshal(msg.Body, &req)
	if err != nil {
		return err
	}
	//解码后对url进行处理todo

	//将处理后的req打包成json返回给前端
	resp, err := p.Marshal()
	if err != nil {
		return fmt.Errorf("marshal response failed: %w", err)
	}
	//生成url ToDo
	_, err = http.Post("115.190.108.142:8080/api/pay/card_pay",
		"application/json",
		bytes.NewBuffer(resp))
	return err
}
