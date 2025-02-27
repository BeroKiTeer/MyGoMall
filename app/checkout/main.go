package main

import (
	"checkout/biz/dal"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/checkout/checkoutservice"
	"github.com/BeroKiTeer/MyGoMall/common/mtl"
	"github.com/BeroKiTeer/MyGoMall/common/serversuite"
	"log"
	"net"
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

	RabbitMQInit()
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
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}

func RabbitMQInit() {
	config, err := conf.GetMQConfig("creditCard")
	if err != nil {
		log.Fatalf("获取支付配置失败: %v", err)
	}
	mqConfig := mq.MQConfig{
		Exchange:     config.Exchange,
		Queue:        config.Queue,
		RoutineKey:   config.RoutingKey,
		ExchangeType: config.ExchangeType,
	}
	mq.CardPaymentProducer, err = mq.NewPaymentProducer(mqConfig)
	config, err = conf.GetMQConfig("creditCard")
	if err != nil {
		log.Fatalf("获取支付配置失败: %v", err)
	}
	mqConfig = mq.MQConfig{
		Exchange:     config.Exchange,
		Queue:        config.Queue,
		RoutineKey:   config.RoutingKey,
		ExchangeType: config.ExchangeType,
	}
	mq.CardPaymentProducer, err = mq.NewPaymentProducer(mqConfig)
}
