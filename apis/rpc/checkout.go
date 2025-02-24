package rpc

import (
	"checkout/conf"
	"github.com/BeroKiTeer/MyGoMall/common/clientsuite"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/checkout/checkoutservice"

	"github.com/cloudwego/kitex/client"
)

var (
	CheckoutClient checkoutservice.Client
)

func initCheckoutClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
		}),
	}
	CheckoutClient, err = checkoutservice.NewClient("checkout", opts...)
	if err != nil {
		panic(err)
	}
}
