package rabbitMq

import (
	"checkout/conf"
	"context"
	"encoding/json"
	"errors"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/stock"
	order_rpc "github.com/BeroKiTeer/MyGoMall/common/rpc/order"
	stock_rpc "github.com/BeroKiTeer/MyGoMall/common/rpc/stock"
	"github.com/streadway/amqp"
)

type PaymentResp struct {
	OrderId       string `json:"order_id"`
	TransactionId string `json:"transaction_id"`
	State         bool   `json:"state"` //支付状态（成功还是失败）
}
type PaymentHandler struct {
	resp PaymentResp
}

func (p *PaymentHandler) GetQueueName() (string, error) {
	config, err := conf.GetQueueConfig("payment_processor")
	if err != nil {
		return "", err
	}
	return config.Queue, nil
}

func (p *PaymentHandler) Unmarshal(data []byte, resp interface{}) error {
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return err
	}
	return nil
}
func (p *PaymentHandler) ProcessMessage(ctx context.Context, msg amqp.Delivery) error {
	var resp PaymentResp

	err := p.Unmarshal(msg.Body, &resp)
	if err != nil {
		return err
	}
	if resp.State == false {
		return errors.New("支付失败！")
	}
	orderClient := order_rpc.DefaultClient()
	stockClient := stock_rpc.DefaultClient()
	//更新订单状态
	_, err = orderClient.MarkOrderPaid(ctx, &order.MarkOrderPaidReq{
		OrderId: resp.OrderId,
	})
	if err != nil {
		return err
	}
	//减少库存
	_, err = stockClient.ReduceItem(ctx, &stock.ReduceItemReq{
		OrderId: resp.OrderId,
	})
	if err != nil {
		return err
	}
	return nil
}
