package rpc

import (
	"checkout/conf"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/checkout/checkoutservice"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	CheckoutClient checkoutservice.Client
)

func initCheckoutClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	CheckoutClient, err = checkoutservice.NewClient("checkout", opts...)
	if err != nil {
		panic(err)
	}
}
