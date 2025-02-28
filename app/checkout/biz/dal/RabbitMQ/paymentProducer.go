package rabbitMq

import (
	"checkout/conf"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

// 支付接口抽象
type PaymentRequest interface {
	GetRoutingKey() string
	Marshal() ([]byte, error)
}

// 生产者基础结构体
type PaymentProducer struct {
	mq       *RabbitMQ
	config   MQConfig
	initOnce sync.Once
}

//初始化生产者

func NewPaymentProducer(config MQConfig) (*PaymentProducer, error) {
	producer := &PaymentProducer{
		mq:     NewRabbitMQ(config.QueueName, config.Exchange, config.RoutingKey),
		config: config,
	}

	// 预先建立连接
	var err error
	producer.mq.Conn, err = amqp.Dial(conf.GetConf().RabbitMQ.MqURL)
	if err != nil {
		return nil, errors.New("建立连接失败！")
	}

	return producer, nil
}

// 初始化AMQP资源（线程安全）
func (p *PaymentProducer) initialize() error {
	var initErr error

	// 1. 声明交换机
	if err := p.mq.Channel.ExchangeDeclare(
		p.config.Exchange,
		p.config.ExchangeType,
		true,  // durable
		false, // autoDelete
		false, // internal
		false, // noWait
		nil,   // args
	); err != nil {
		initErr = fmt.Errorf("声明交换机失败: %w", err)
		return err
	}

	// 2. 声明队列
	if _, err := p.mq.Channel.QueueDeclare(
		p.config.QueueName,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	); err != nil {
		initErr = fmt.Errorf("声明队列失败: %w", err)
		return err
	}

	// 3. 绑定队列
	if err := p.mq.Channel.QueueBind(
		p.config.QueueName,
		p.config.RoutingKey,
		p.config.Exchange,
		false, // noWait
		nil,   // args
	); err != nil {
		initErr = fmt.Errorf("绑定队列失败: %w", err)
	}

	return initErr
}

// 发送支付请求
func (p *PaymentProducer) Send(req PaymentRequest) error {
	// 1. 确保初始化完成
	p.initOnce.Do(func() {
		err := p.initialize()
		if err != nil {
			return
		} // 保证只执行一次
	})
	//有一个问题，就是这段代码可以让生产者与多个队列建立连接吗？
	// 2. 序列化消息
	data, err := req.Marshal()
	if err != nil {
		return fmt.Errorf("序列化失败: %w", err)
	}

	// 3. 发送消息
	return p.mq.Channel.Publish(
		p.config.Exchange,
		req.GetRoutingKey(), // 使用动态路由键
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp.Persistent, // 持久化消息
			Timestamp:    time.Now(),
		},
	)
}
