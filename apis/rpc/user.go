package rpc

import (
	"apis/conf"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"user/kitex_gen/user/userservice"
)

var (
	UserClient userservice.Client
)

func initUserClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	UserClient, err = userservice.NewClient("user", opts...)
	if err != nil {
		panic(err)
	}
}
