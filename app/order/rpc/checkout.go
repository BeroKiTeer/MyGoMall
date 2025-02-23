package rpc

import (
	"cart/conf"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	CheckoutClient checkoutservice.Client
)

func initCheckoutClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		klog.Fatal(err)
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	CheckoutClient, err = checkoutservice.NewClient("checkout", opts...)
	if err != nil {
		klog.Fatal(err)
		panic(err)
	}
}
