package rpc

import (
	"auth/conf"
	"github.com/BeroKiTeer/MyGoMall/common/clientsuite"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/auth/authservice"
	"github.com/cloudwego/kitex/client"
)

var (
	AuthClient   authservice.Client
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
	err          error
)

func initAuthClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}
	AuthClient, err = authservice.NewClient("auth", opts...)
	if err != nil {
		panic(err)
	}
}
