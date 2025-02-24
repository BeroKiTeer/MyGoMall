package rpc

import (
	"github.com/BeroKiTeer/MyGoMall/common/clientsuite"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/client"
	"order/conf"
)

var (
	OrderClient orderservice.Client
)

func initOrderClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
		}),
	}
	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		panic(err)
	}
}
