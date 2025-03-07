package rpc

import (
	"apis/conf"
	"github.com/BeroKiTeer/MyGoMall/common/clientsuite"
	"github.com/BeroKiTeer/MyGoMall/common/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
)

var (
	UserClient userservice.Client
)

func initUserClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonClientSuite{
			CurrentServiceName: conf.GetConf().Hertz.Service,
			RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
		}),
	}
	UserClient, err = userservice.NewClient("user", opts...)
	if err != nil {
		panic(err)
	}
}
