package rpc

import (
	"github.com/BeroKiTeer/MyGoMall/common/clientsuite"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"product/conf"
)

var (
	ProductClient productcatalogservice.Client
)

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		panic(err)
	}
}
