package stock

import (
	"github.com/cloudwego/kitex/pkg/discovery"
	"sync"

	"github.com/cloudwego/kitex/client"
)

var (
	// todo edit custom config
	defaultClient     RPCClient
	defaultDstService = "stock"
	consulResolver    discovery.Resolver // 解析器字段
	once              sync.Once
)

func init() {
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
