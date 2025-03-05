package main

import (
	"checkout/biz/dal"
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/checkout/checkoutservice"
	"github.com/BeroKiTeer/MyGoMall/common/mtl"
	"github.com/BeroKiTeer/MyGoMall/common/serversuite"
	"log"
	"net"
	"os"
	"time"

	mq "checkout/biz/dal/RabbitMQ"
	"checkout/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func main() {
	dal.Init()
	mtl.InitMetric(ServiceName, conf.GetConf().Kitex.MetricsPort, RegistryAddr)
	mtl.InitTracing(ServiceName)
	opts := kitexInit()
	go PaymentConsumerInit()
	PymentProducerInit()
	svr := checkoutservice.NewServer(new(CheckoutServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr), server.WithSuite(serversuite.CommonServerSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	consoleOutput := zapcore.Lock(os.Stderr) // 线程安全控制台输出
	multiOutput := zapcore.NewMultiWriteSyncer(asyncWriter, consoleOutput)
	klog.SetOutput(multiOutput)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}

func PymentProducerInit() {

	config, err := conf.GetMQConfig("creditCard")
	if err != nil {
		log.Fatalf("获取支付配置失败: %v", err)
	}
	log.Printf("尝试连接RabbitMQ: %s", config.URL)

	mqConfig := mq.MQConfig{
		Exchange:     config.Exchange,
		QueueName:    config.Queue,
		RoutingKey:   config.RoutingKey,
		ExchangeType: config.ExchangeType,
	}
	mq.CardPaymentProducer, err = mq.NewPaymentProducer(mqConfig)
}

func PaymentConsumerInit() {
	// 1. 获取MQ配置
	config := conf.GetConf().RabbitMQ.Consumers.Processors["payment_processor"]

	// 2. 创建消费者实例
	consumer, err := mq.GetCheckoutConsumer(&mq.MQConfig{
		Exchange:     config.Exchange,
		QueueName:    config.Queue,
		ExchangeType: config.ExchangeType,
	})
	if err != nil {
		log.Fatalf("创建消费者失败: %v", err)
	}

	// 3. 绑定队列
	for _, key := range config.BindingKeys {
		if err := consumer.BindQueue(config.Queue, config.Exchange, key); err != nil {
			log.Fatalf("队列绑定失败: %v", err)
		}
	}

	// 4. 启动消费监听
	ctx := context.Background()
	handler := &mq.PaymentHandler{}
	if err := consumer.Consume(ctx, handler); err != nil {
		klog.Errorf("消费异常终止: %v", err)
	}
}
