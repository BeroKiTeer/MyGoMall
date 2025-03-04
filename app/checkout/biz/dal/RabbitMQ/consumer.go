package rabbitMq

import (
	"checkout/conf"
	"context"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
)

type MessageHandler interface {
	Unmarshal([]byte, interface{}) error
	GetQueueName() (string, error)
	ProcessMessage(ctx context.Context, msg amqp.Delivery) error //对消息进行处理
}

type CheckoutConsumer struct {
	mq       *RabbitMQ
	handlers []MessageHandler
	config   *MQConfig
}

func GetCheckoutConsumer(config *MQConfig) (*CheckoutConsumer, error) {
	consumer := &CheckoutConsumer{
		mq:       NewRabbitMQ("", "", ""),
		handlers: nil,
		config:   config,
	}
	// 预先建立连接
	var err error
	consumer.mq.Conn, err = amqp.Dial(conf.GetConf().RabbitMQ.MqURL)
	if err != nil {
		return nil, errors.New("建立连接失败！")
	}

	return consumer, nil
}

func (p *CheckoutConsumer) BindQueue(QueueName, Exchange, RoutingKey string) error {
	_, err := p.mq.Channel.QueueDeclare( // 返回的队列对象内部记录了队列的一些信息，这里没什么用
		QueueName, // 队列名
		true,      // 是否持久化
		false,     // 是否自动删除(前提是至少有一个消费者连接到这个队列，之后所有与这个队列连接的消费者都断开时，才会自动删除。注意：生产者客户端创建这个队列，或者没有消费者客户端与这个队列连接时，都不会自动删除这个队列)
		false,     // 是否为排他队列（排他的队列仅对“首次”声明的conn可见[一个conn中的其他channel也能访问该队列]，conn结束后队列删除）
		false,     // 是否阻塞
		nil,       // 额外属性（我还不会用）
	)
	if err != nil {
		fmt.Println("声明队列失败", err)
		return err
	}
	//绑定队列到交换机
	err = p.mq.Channel.QueueBind(
		QueueName,  // 队列名
		RoutingKey, // 绑定键（支持通配符）
		Exchange,   // 交换机名
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("队列绑定失败: %w", err)
	}
	return nil
}

// 消费者指定一条队列消费信息
func (p *CheckoutConsumer) Consume(ctx context.Context, msgHandler MessageHandler) error {
	// 2.从队列获取消息（消费者只关注队列）consume方式会不断的从队列中获取消息
	QueueName, err := msgHandler.GetQueueName()
	if err != nil {
		return err
	}
	msgChanl, err := p.mq.Channel.Consume(
		QueueName, // 队列名
		"",        // 消费者名，用来区分多个消费者，以实现公平分发或均等分发策略
		true,      // 是否自动应答
		false,     // 是否排他
		false,     // 是否接收只同一个连接中的消息，若为true，则只能接收别的conn中发送的消息
		true,      // 队列消费是否阻塞
		nil,       // 额外属性
	)
	if err != nil {
		fmt.Println("获取消息失败", err)
		return err
	}
	for msg := range msgChanl {
		if err := msgHandler.ProcessMessage(ctx, msg); err != nil {
			// 处理失败逻辑
			//msg.Nack(false, shouldRetry(err))
			err := msg.Nack(false, true)
			if err != nil {
				return err
			} // 重试
			return err
		} else {
			err := msg.Ack(false)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
