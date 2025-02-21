package rpc

import (
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
	"user/conf"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func InitClient() {
	once.Do(func() {
		initAuthClient()
	})
}

func initAuthClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		panic(err)
	}
}
