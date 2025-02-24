package rpc

import (
	"github.com/BeroKiTeer/MyGoMall/common/clientsuite"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/payment/paymentservice"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"product/conf"
)

var (
	PaymentClient paymentservice.Client
)

func initPaymentClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
		}),
	}
	ProductClient, err = productcatalogservice.NewClient("payment", opts...)
	if err != nil {
		panic(err)
	}
}
