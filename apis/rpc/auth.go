package rpc

import (
	"apis/conf"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/auth/authservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	AuthClient authservice.Client
)

func initAuthClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	AuthClient, err = authservice.NewClient("auth", opts...)
	if err != nil {
		panic(err)
	}
}
