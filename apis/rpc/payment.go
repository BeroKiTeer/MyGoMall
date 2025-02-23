package rpc

import (
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/payment/paymentservice"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"product/conf"
)

var (
	PaymentClient paymentservice.Client
)

func initPaymentClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("payment", opts...)
	if err != nil {
		panic(err)
	}
}
