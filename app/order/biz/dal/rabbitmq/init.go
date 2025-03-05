package rabbitmq

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"order/conf"

	"github.com/streadway/amqp"
) //导入mq包

// MQConfig MqURL 格式 amqp://账号：密码@rabbitmq服务器地址：端口号/vhost (默认是5672端口)
// 端口可在 /etc/rabbitmq/rabbitmq-env.conf 配置文件设置，也可以启动后通过netstat -tlnp查看
type MQConfig struct {
	Exchange     string
	QueueName    string
	RoutingKey   string
	ExchangeType string
	MqURL        string
}

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	config  *MQConfig
}

// NewRabbitMQ 创建结构体实例
func NewRabbitMQ(queueName, exchange, routingKey string) *RabbitMQ {
	rabbitMQ := RabbitMQ{config: &MQConfig{
		Exchange:     exchange,
		QueueName:    queueName,
		RoutingKey:   routingKey,
		ExchangeType: "Topic",
		MqURL:        conf.GetConf().RabbitMQ.MqURL,
	}}
	var err error
	//创建rabbitmq连接
	rabbitMQ.Conn, err = amqp.Dial(rabbitMQ.config.MqURL)
	checkErr(err, "创建连接失败")

	//创建Channel
	rabbitMQ.Channel, err = rabbitMQ.Conn.Channel()
	checkErr(err, "创建channel失败")

	return &rabbitMQ

}

// ReleaseRes 释放资源,建议NewRabbitMQ获取实例后 配合defer使用
func (mq *RabbitMQ) ReleaseRes() {
	err := mq.Conn.Close()
	if err != nil {
		klog.Errorf("%v", err)
		return
	}
	err = mq.Channel.Close()
	if err != nil {
		klog.Errorf("%v", err)
		return
	}
}

func checkErr(err error, meg string) {
	if err != nil {
		klog.Fatalf("%s:%s\n", meg, err)
	}
}
