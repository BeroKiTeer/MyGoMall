package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"net"
	"payment/conf"
	"time"
)

type MessageHandler interface {
	Marshal() ([]byte, error)
	Unmarshal([]byte, interface{}) error
	GetQueueName() (string, error)
	ProcessMessage(ctx context.Context, msg amqp.Delivery) error //对消息进行处理
}

type PaymentConsumer struct {
	mq       *RabbitMQ
	handlers []MessageHandler
	config   *MQConfig
}

func GetPaymentConsumer(config *MQConfig) (*PaymentConsumer, error) {
	consumer := &PaymentConsumer{
		mq:       NewRabbitMQ("", "", ""),
		handlers: nil,
		config:   config,
	}
	// 预先建立连接
	var err error
	consumer.mq.Conn, err = amqp.Dial(conf.GetConf().RabbitMQ.RabbitMQURL)
	if err != nil {
		return nil, errors.New("建立连接失败！")
	}

	return consumer, nil
}

func (p *PaymentConsumer) BindQueue(QueueName, Exchange, RoutingKey string) error {
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
func (p *PaymentConsumer) Consume(ctx context.Context, msgHandler MessageHandler) error {
	retryCount := 0
	maxRetries := 10 // 最大重试次数

	for {
		select {
		case <-ctx.Done():
			klog.Info("接收到终止信号，停止消费")
			return nil
		default:
			QueueName, err := msgHandler.GetQueueName()
			if err != nil {
				return fmt.Errorf("获取队列名失败: %w", err)
			}

			// 重建连接（关键）
			if err := p.reconnectIfNeeded(); err != nil {
				klog.Errorf("连接重建失败: %v", err)
				if retryCount++; retryCount > maxRetries {
					return fmt.Errorf("超过最大重试次数")
				}
				time.Sleep(3 * time.Second)
				continue
			}

			// 创建消费通道
			msgChan, err := p.mq.Channel.Consume(
				QueueName,
				"consumer_"+uuid.New().String(), // 唯一消费者标识
				false,                           // 关闭自动ACK
				false,
				false,
				false,
				nil,
			)
			if err != nil {
				klog.Errorf("创建消费通道失败: %v", err)
				time.Sleep(3 * time.Second)
				continue
			}

			// 重置重试计数器
			retryCount = 0

			// 消息处理循环
			for msg := range msgChan {
				if err := p.handleMessage(msgHandler, msg); err != nil {
					klog.Errorf("消息处理失败: %v", err)
				}
			}

			klog.Warn("消息通道关闭，等待重连...")
			time.Sleep(5 * time.Second)
		}
	}
}

// 消息处理封装
func (p *PaymentConsumer) handleMessage(h MessageHandler, msg amqp.Delivery) error {
	defer func() {
		if r := recover(); r != nil {
			klog.Errorf("消息处理发生panic: %v", r)
			msg.Nack(false, true) // 要求重试
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := h.ProcessMessage(ctx, msg); err != nil {
		if shouldRetry(err) {
			msg.Nack(false, true) // 重试消息
			return err
		}
		msg.Nack(false, false) // 丢弃消息
		return fmt.Errorf("不可恢复错误: %w", err)
	}

	msg.Ack(false)
	return nil
}

// 消息处理封装
func (p *PaymentConsumer) reconnectIfNeeded() error {
	// 检查连接状态
	if p.mq.Conn == nil || p.mq.Conn.IsClosed() {
		newConn, err := amqp.Dial(conf.GetConf().RabbitMQ.RabbitMQURL)
		if err != nil {
			return fmt.Errorf("连接重建失败: %w", err)
		}
		p.mq.Conn = newConn
	}

	// 通道有效性检测（新增核心逻辑）
	if p.mq.Channel != nil {
		// 通过无害操作检测通道状态
		_, err := p.mq.Channel.QueueDeclarePassive(
			"amq.rabbitmq.reply-to", // RabbitMQ内置队列
			false,                   // 非持久化
			true,                    // 自动删除
			true,                    // 排他队列
			false,
			nil,
		)
		if err != nil {
			p.mq.Channel = nil // 标记通道失效
		}
	}

	// 重建通道
	if p.mq.Channel == nil {
		ch, err := p.mq.Conn.Channel()
		if err != nil {
			return fmt.Errorf("通道重建失败: %w", err)
		}
		p.mq.Channel = ch

		// 设置通道关闭监听（关键增强）
		closeChan := make(chan *amqp.Error)
		p.mq.Channel.NotifyClose(closeChan)
		go func() {
			<-closeChan
			p.mq.Channel = nil // 通道关闭时自动置空
			klog.Warn("检测到通道异常关闭")
		}()
	}

	return nil
}

func shouldRetry(err error) bool {
	// 示例逻辑：网络类错误或超时错误允许重试
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		return true
	}
	// 可扩展其他重试条件
	return false
}
