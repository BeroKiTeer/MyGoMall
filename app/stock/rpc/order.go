package rpc

import (
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	consul "github.com/kitex-contrib/registry-consul"
	"order/conf"
)

var (
	OrderClient orderservice.Client
)

func initOrderClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		klog.Fatal(err)
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		klog.Fatal(err)
		panic(err)
	}
}
