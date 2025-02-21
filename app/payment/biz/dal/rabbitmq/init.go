package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"payment/conf"
	"time"
)

var (
	rabbitMQURL        = conf.GetConf().RabbitMQ.RabbitMQURL
	paymentExchange    = conf.GetConf().RabbitMQ.PaymentExchange
	paymentQueue       = conf.GetConf().RabbitMQ.PaymentQueue
	paymentDLXExchange = conf.GetConf().RabbitMQ.PaymentDLXExchange
	paymentDLXQueue    = conf.GetConf().RabbitMQ.PaymentDLXQueue
)

// 连接到 RabbitMQ
func connectToRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		return nil, nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		err := conn.Close()
		if err != nil {
			return nil, nil, err
		}
		return nil, nil, err
	}
	return conn, ch, nil
}

// 声明交换器和队列
func declareExchangesAndQueues(ch *amqp.Channel) error {
	// 声明死信交换器
	err := ch.ExchangeDeclare(
		paymentDLXExchange, // 交换器名称
		"direct",           // 交换器类型
		true,               // 是否持久化
		false,              // 是否自动删除
		false,              // 是否内部使用
		false,              // 是否等待服务器响应
		nil,                // 额外参数
	)
	if err != nil {
		return err
	}

	// 声明死信队列
	_, err = ch.QueueDeclare(
		paymentDLXQueue, // 队列名称
		true,            // 是否持久化
		false,           // 是否自动删除
		false,           // 是否排他
		false,           // 是否等待服务器响应
		nil,             // 额外参数
	)
	if err != nil {
		return err
	}

	// 绑定死信队列到死信交换器
	err = ch.QueueBind(
		paymentDLXQueue,    // 队列名称
		"payment_dlx_key",  // 路由键
		paymentDLXExchange, // 交换器名称
		false,              // 是否等待服务器响应
		nil,                // 额外参数
	)
	if err != nil {
		return err
	}

	// 声明主交换器
	err = ch.ExchangeDeclare(
		paymentExchange, // 交换器名称
		"direct",        // 交换器类型
		true,            // 是否持久化
		false,           // 是否自动删除
		false,           // 是否内部使用
		false,           // 是否等待服务器响应
		nil,             // 额外参数
	)
	if err != nil {
		return err
	}

	// 声明主队列，并设置死信交换器和路由键
	args := amqp.Table{
		"x-dead-letter-exchange":    paymentDLXExchange,
		"x-dead-letter-routing-key": "payment_dlx_key",
	}
	_, err = ch.QueueDeclare(
		paymentQueue, // 队列名称
		true,         // 是否持久化
		false,        // 是否自动删除
		false,        // 是否排他
		false,        // 是否等待服务器响应
		args,         // 额外参数
	)
	if err != nil {
		return err
	}

	// 绑定主队列到主交换器
	err = ch.QueueBind(
		paymentQueue,    // 队列名称
		"payment_key",   // 路由键
		paymentExchange, // 交换器名称
		false,           // 是否等待服务器响应
		nil,             // 额外参数
	)
	if err != nil {
		return err
	}

	return nil
}

// 发送订单消息到延迟队列
func sendOrderToQueue(ID string, delay time.Duration) error {
	conn, ch, err := connectToRabbitMQ()
	if err != nil {
		return err
	}
	defer conn.Close()
	defer ch.Close()

	err = declareExchangesAndQueues(ch)
	if err != nil {
		return err
	}

	// 设置消息的 TTL
	expiration := fmt.Sprintf("%d", delay.Milliseconds())
	msg := amqp.Publishing{
		ContentType:  "text/plain",
		Body:         []byte(ID),
		Expiration:   expiration,
		DeliveryMode: amqp.Persistent,
	}

	// 发布消息到主交换器
	err = ch.Publish(
		paymentExchange, // 交换器名称
		"payment_key",   // 路由键
		false,           // 是否强制
		false,           // 是否立即
		msg,
	)
	if err != nil {
		return err
	}

	log.Printf("Order %s sent to queue with delay %v", ID, delay)
	return nil
}

// 从死信队列消费消息并取消支付
func consumeFromDLXQueue() error {
	conn, ch, err := connectToRabbitMQ()
	if err != nil {
		return err
	}
	defer conn.Close()
	defer ch.Close()

	err = declareExchangesAndQueues(ch)
	if err != nil {
		return err
	}

	// 消费死信队列中的消息
	msgs, err := ch.Consume(
		paymentDLXQueue, // 队列名称
		"",              // 消费者名称
		true,            // 是否自动确认
		false,           // 是否排他
		false,           // 是否为本地队列
		false,           // 是否等待服务器响应
		nil,             // 额外参数
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			orderID := string(d.Body)
			log.Printf("Received order %s from DLX queue, canceling payment...", orderID)
			// test为rabbitmq初始化
			if orderID == "test" {
				log.Printf("rabbitMQ初始化成功...")
			}
			//调用取消支付的业务逻辑

		}
	}()

	log.Printf(" [*] Waiting for messages from DLX queue. To exit press CTRL+C")
	<-forever

	return nil
}

func Init() {
	err := sendOrderToQueue("test", 1)
	if err != nil {
		log.Fatal(err)
	}
	err = consumeFromDLXQueue()
	if err != nil {
		log.Fatal(err)
	}
}
