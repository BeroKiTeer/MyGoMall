package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/BeroKiTeer/MyGoMall/common/constant"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
	"order/biz/dal/mysql"
	"order/biz/model"
	"order/conf"
)

type Message struct {
	OrderId       string `json:"order_id"`
	TransactionId string `json:"transaction_id"`
	Success       bool   `json:"success"` //支付状态（成功还是失败）
}
type PaymentHandler struct {
	msg Message
}

func (p *PaymentHandler) GetQueueName() (string, error) {
	config, err := conf.GetQueueConfig("payment_processor")
	if err != nil {
		klog.Errorf("get queue config error: %v", err)
		return "", err
	}
	return config.Queue, nil
}

func (p *PaymentHandler) Unmarshal(data []byte, resp interface{}) error {
	err := json.Unmarshal(data, &resp)
	if err != nil {
		klog.Errorf("unmarshal message error: %v", err)
		return err
	}
	return nil
}
func (p *PaymentHandler) ProcessMessage(ctx context.Context, msg amqp.Delivery) error {
	var resp Message
	klog.Infof("收到消息: %v", string(msg.Body))
	err := p.Unmarshal(msg.Body, &resp)
	if err != nil {
		return err
	}
	if resp.Success == false {
		klog.Errorf("支付失败！")
		return errors.New("支付失败！")
	}
	// Finish your business logic.
	// 参数验证
	if resp.OrderId == "" {
		klog.Errorf("没有OrderID")
		return err
	}
	// 获取订单中的商品信息
	//productIds := model.GetProductIdsFromOrder(mysql.DB, resp.OrderId)
	//
	//products := make([]*product.GetProductResp, 1)
	//for _, productID := range productIds {
	//	res, err := rpc.ProductClient.GetProduct(ctx, &product.GetProductReq{Id: productID})
	//	if err != nil {
	//		klog.Errorf("get product error: %v", err)
	//	}
	//	products = append(products, res)
	//}

	//_, err = rpc.StockClient.ReduceItem(ctx, &stock.ReduceItemReq{
	//	OrderId: resp.OrderId,
	//})
	//if err != nil {
	//	klog.Errorf("%v", err)
	//	return err
	//}

	// 2. 订单状态改为已支付
	if err = model.UpdateOrderStatus(mysql.DB, resp.OrderId, constant.Paid); err != nil {
		klog.Errorf("%v", err)
		return err
	}
	klog.Infof("订单状态改为已支付")
	return nil
}
