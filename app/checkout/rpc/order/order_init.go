package order

import (
	"checkout/conf"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	// todo edit custom config
	defaultClient     RPCClient
	defaultDstService = "order"
	consulResolver    discovery.Resolver // 解析器字段
	once              sync.Once
)

func init() {
	var err error
	consulResolver, err = consul.NewConsulResolver(
		conf.GetConf().Registry.RegistryAddress[0],
	)
	if err != nil {
		panic("failed to create consul resolver: " + err.Error())
	}
	DefaultClient()
}

func DefaultClient() RPCClient {
	once.Do(func() {
		opts := []client.Option{
			client.WithResolver(consulResolver), // 使用Consul解析器
		}
		defaultClient = newClient(defaultDstService, opts...)
	})
	return defaultClient
}

func newClient(dstService string, opts ...client.Option) RPCClient {
	c, err := NewRPCClient(dstService, opts...)
	if err != nil {
		panic("failed to init client: " + err.Error())
	}
	return c
}

func InitClient(dstService string, opts ...client.Option) {
	baseOpts := []client.Option{
		client.WithResolver(consulResolver),
	}
	opts = append(baseOpts, opts...)

	defaultClient = newClient(dstService, opts...)
}
