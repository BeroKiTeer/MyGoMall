package clientsuite

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonClientSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

func (s CommonClientSuite) Options() []client.Option {
	opts := []client.Option{
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName, // 当前服务名称
		}),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler), // 添加HTTP2元数据处理中间件
		client.WithTransportProtocol(transport.GRPC),         // 指定GRPC传输协议
		client.WithSuite(tracing.NewClientSuite()),           // 集成OpenTelemetry追踪套件
	}
	// 创建Consul服务发现解析器
	r, err := consul.NewConsulResolver(s.RegistryAddr) // 从配置获取注册中心地址
	if err != nil {
		panic(err)
	} // 服务发现初始化失败时终止进程
	opts = append(opts, client.WithResolver(r)) // 将服务发现组件加入配置选项
	return opts
}
