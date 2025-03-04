package main

import (
	"context"
	"github.com/BeroKiTeer/MyGoMall/common/mtl"
	"github.com/BeroKiTeer/MyGoMall/common/serversuite"
	"log"
	"net"
	"os"
	"payment/biz/dal"
	mq "payment/biz/dal/rabbitmq"
	"time"

	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"payment/conf"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
)

func main() {
	mtl.InitMetric(ServiceName, conf.GetConf().Kitex.MetricsPort, RegistryAddr)
	mtl.InitTracing(ServiceName)
	dal.Init()
	PaymentConsumerInit()
	opts := kitexInit()

	svr := paymentservice.NewServer(new(PaymentServiceImpl), opts...)

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

func PaymentConsumerInit() {
	// 1. 获取MQ配置
	config := conf.GetConf().RabbitMQ.Consumers.Processors["payment_processor"]

	// 2. 创建消费者实例
	consumer, err := mq.GetPaymentConsumer(&mq.MQConfig{
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
		log.Printf("消费异常终止: %v", err)
	}
}
