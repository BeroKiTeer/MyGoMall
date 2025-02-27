package rpc

import "sync"

var (
	once sync.Once
	err  error
)

func InitClient() {
	once.Do(func() {
		initAuthClient()
		initUserClient()
		initProductClient()
		initCartClient()
		initOrderClient()
		//initCheckoutClient()
		initPaymentClient()
	})
}
