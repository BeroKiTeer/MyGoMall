package rpc

import (
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"product/conf"
	"product/kitex_gen/product/productcatalogservice"
)

var (
	ProductClient productcatalogservice.Client
)

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("cart", opts...)
	if err != nil {
		panic(err)
	}
}
