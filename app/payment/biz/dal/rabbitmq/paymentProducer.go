package rabbitmq

import (
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/streadway/amqp"
	"payment/conf"
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
	mq          *RabbitMQ
	config      MQConfig
	initOnce    sync.Once
	channelPool sync.Pool
}

//初始化生产者

func NewPaymentProducer(config MQConfig) (*PaymentProducer, error) {
	producer := &PaymentProducer{
		mq:     NewRabbitMQ(config.QueueName, config.Exchange, config.RoutingKey),
		config: config,
	}

	// 预先建立连接
	var err error
	producer.mq.Conn, err = amqp.Dial(conf.GetConf().RabbitMQ.RabbitMQURL)
	if err != nil {
		return nil, errors.New("建立连接失败！")
	}
	klog.Info("连接成功")
	producer.initOnce.Do(func() {
		err := producer.initialize()
		if err != nil {
			return
		} // 保证只执行一次
	})
	producer.channelPool = sync.Pool{
		New: func() interface{} {
			ch, err := producer.mq.Conn.Channel()
			if err != nil {
				return nil
			}
			ch.Confirm(false)
			return ch
		},
	}
	return producer, nil
}

// 初始化AMQP资源（线程安全）
func (p *PaymentProducer) initialize() error {
	var initErr error

	if p.config.Exchange == "" {
		return errors.New("交换机名称不能为空")
	}

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

// 发送支付结果
func (p *PaymentProducer) Send(req PaymentRequest) error {
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

// 新增通道检查方法
func (p *PaymentProducer) ensureChannel() error {
	// 修改判断条件
	if p.mq.Channel == nil {
		ch, err := p.mq.Conn.Channel()
		if err != nil {
			return fmt.Errorf("通道创建失败: %w", err)
		}
		p.mq.Channel = ch

		// 增加通道关闭监听
		closeChan := ch.NotifyClose(make(chan *amqp.Error))
		go func() {
			<-closeChan
			p.mq.Channel = nil // 通道关闭时主动置空
			klog.Warn("AMQP通道异常关闭")
		}()

		if err := ch.Confirm(false); err != nil {
			return fmt.Errorf("启用确认模式失败: %w", err)
		}
	}
	return nil
}
